package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetRoutersCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetRoutersCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["routersCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetRoutersCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRouterCountViaNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetRouterCountViaName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["routerCount"] = 0
	}
	m["routerCount"] = _int64
	return self.JSON(m)
}

func GetRouterCountViaPathHandler(self *macross.Context) error {
	Path_ := self.Args("path").String()
	_int64 := model.GetRouterCountViaPath(Path_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["routerCount"] = 0
	}
	m["routerCount"] = _int64
	return self.JSON(m)
}

func GetRouterCountViaPatternOutlineHandler(self *macross.Context) error {
	PatternOutline_ := self.Args("pattern_outline").String()
	_int64 := model.GetRouterCountViaPatternOutline(PatternOutline_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["routerCount"] = 0
	}
	m["routerCount"] = _int64
	return self.JSON(m)
}

func GetRouterCountViaFitHandler(self *macross.Context) error {
	Fit_ := self.Args("fit").MustInt()
	_int64 := model.GetRouterCountViaFit(Fit_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["routerCount"] = 0
	}
	m["routerCount"] = _int64
	return self.JSON(m)
}

func GetRouterCountViaRouteHandler(self *macross.Context) error {
	Route_ := self.Args("route").Bytes()
	_int64 := model.GetRouterCountViaRoute(Route_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["routerCount"] = 0
	}
	m["routerCount"] = _int64
	return self.JSON(m)
}

func GetRouterCountViaNumberPartsHandler(self *macross.Context) error {
	NumberParts_ := self.Args("number_parts").MustInt()
	_int64 := model.GetRouterCountViaNumberParts(NumberParts_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["routerCount"] = 0
	}
	m["routerCount"] = _int64
	return self.JSON(m)
}

func GetRoutersViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_Router, _error := model.GetRoutersViaName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPath := self.Args("path").String()
	if (offset > 0) && helper.IsHas(iPath) {
		_Router, _error := model.GetRoutersViaPath(offset, limit, iPath, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersViaPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersViaPatternOutlineHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPatternOutline := self.Args("pattern_outline").String()
	if (offset > 0) && helper.IsHas(iPatternOutline) {
		_Router, _error := model.GetRoutersViaPatternOutline(offset, limit, iPatternOutline, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersViaPatternOutline's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersViaFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFit := self.Args("fit").MustInt()
	if (offset > 0) && helper.IsHas(iFit) {
		_Router, _error := model.GetRoutersViaFit(offset, limit, iFit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersViaFit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersViaRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRoute := self.Args("route").Bytes()
	if (offset > 0) && helper.IsHas(iRoute) {
		_Router, _error := model.GetRoutersViaRoute(offset, limit, iRoute, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersViaRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersViaNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iNumberParts := self.Args("number_parts").MustInt()
	if (offset > 0) && helper.IsHas(iNumberParts) {
		_Router, _error := model.GetRoutersViaNumberParts(offset, limit, iNumberParts, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersViaNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndPathAndPatternOutlineHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPath := self.Args("path").String()
	iPatternOutline := self.Args("pattern_outline").String()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndPathAndPatternOutline(offset, limit, iName,iPath,iPatternOutline)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndPathAndPatternOutline's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndPathAndFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPath := self.Args("path").String()
	iFit := self.Args("fit").MustInt()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndPathAndFit(offset, limit, iName,iPath,iFit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndPathAndFit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndPathAndRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPath := self.Args("path").String()
	iRoute := self.Args("route").Bytes()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndPathAndRoute(offset, limit, iName,iPath,iRoute)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndPathAndRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndPathAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPath := self.Args("path").String()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndPathAndNumberParts(offset, limit, iName,iPath,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndPathAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndPatternOutlineAndFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPatternOutline := self.Args("pattern_outline").String()
	iFit := self.Args("fit").MustInt()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndPatternOutlineAndFit(offset, limit, iName,iPatternOutline,iFit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndPatternOutlineAndFit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndPatternOutlineAndRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPatternOutline := self.Args("pattern_outline").String()
	iRoute := self.Args("route").Bytes()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndPatternOutlineAndRoute(offset, limit, iName,iPatternOutline,iRoute)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndPatternOutlineAndRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndPatternOutlineAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPatternOutline := self.Args("pattern_outline").String()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndPatternOutlineAndNumberParts(offset, limit, iName,iPatternOutline,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndPatternOutlineAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndFitAndRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iFit := self.Args("fit").MustInt()
	iRoute := self.Args("route").Bytes()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndFitAndRoute(offset, limit, iName,iFit,iRoute)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndFitAndRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndFitAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iFit := self.Args("fit").MustInt()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndFitAndNumberParts(offset, limit, iName,iFit,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndFitAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndRouteAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iRoute := self.Args("route").Bytes()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndRouteAndNumberParts(offset, limit, iName,iRoute,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndRouteAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPathAndPatternOutlineAndFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iPatternOutline := self.Args("pattern_outline").String()
	iFit := self.Args("fit").MustInt()

	if helper.IsHas(iPath) {
		_Router, _error := model.GetRoutersByPathAndPatternOutlineAndFit(offset, limit, iPath,iPatternOutline,iFit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPathAndPatternOutlineAndFit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPathAndPatternOutlineAndRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iPatternOutline := self.Args("pattern_outline").String()
	iRoute := self.Args("route").Bytes()

	if helper.IsHas(iPath) {
		_Router, _error := model.GetRoutersByPathAndPatternOutlineAndRoute(offset, limit, iPath,iPatternOutline,iRoute)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPathAndPatternOutlineAndRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPathAndPatternOutlineAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iPatternOutline := self.Args("pattern_outline").String()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iPath) {
		_Router, _error := model.GetRoutersByPathAndPatternOutlineAndNumberParts(offset, limit, iPath,iPatternOutline,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPathAndPatternOutlineAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPathAndFitAndRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iFit := self.Args("fit").MustInt()
	iRoute := self.Args("route").Bytes()

	if helper.IsHas(iPath) {
		_Router, _error := model.GetRoutersByPathAndFitAndRoute(offset, limit, iPath,iFit,iRoute)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPathAndFitAndRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPathAndFitAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iFit := self.Args("fit").MustInt()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iPath) {
		_Router, _error := model.GetRoutersByPathAndFitAndNumberParts(offset, limit, iPath,iFit,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPathAndFitAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPathAndRouteAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRoute := self.Args("route").Bytes()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iPath) {
		_Router, _error := model.GetRoutersByPathAndRouteAndNumberParts(offset, limit, iPath,iRoute,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPathAndRouteAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPatternOutlineAndFitAndRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPatternOutline := self.Args("pattern_outline").String()
	iFit := self.Args("fit").MustInt()
	iRoute := self.Args("route").Bytes()

	if helper.IsHas(iPatternOutline) {
		_Router, _error := model.GetRoutersByPatternOutlineAndFitAndRoute(offset, limit, iPatternOutline,iFit,iRoute)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPatternOutlineAndFitAndRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPatternOutlineAndFitAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPatternOutline := self.Args("pattern_outline").String()
	iFit := self.Args("fit").MustInt()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iPatternOutline) {
		_Router, _error := model.GetRoutersByPatternOutlineAndFitAndNumberParts(offset, limit, iPatternOutline,iFit,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPatternOutlineAndFitAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPatternOutlineAndRouteAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPatternOutline := self.Args("pattern_outline").String()
	iRoute := self.Args("route").Bytes()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iPatternOutline) {
		_Router, _error := model.GetRoutersByPatternOutlineAndRouteAndNumberParts(offset, limit, iPatternOutline,iRoute,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPatternOutlineAndRouteAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByFitAndRouteAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFit := self.Args("fit").MustInt()
	iRoute := self.Args("route").Bytes()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iFit) {
		_Router, _error := model.GetRoutersByFitAndRouteAndNumberParts(offset, limit, iFit,iRoute,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByFitAndRouteAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPath := self.Args("path").String()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndPath(offset, limit, iName,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndPatternOutlineHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPatternOutline := self.Args("pattern_outline").String()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndPatternOutline(offset, limit, iName,iPatternOutline)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndPatternOutline's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iFit := self.Args("fit").MustInt()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndFit(offset, limit, iName,iFit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndFit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iRoute := self.Args("route").Bytes()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndRoute(offset, limit, iName,iRoute)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByNameAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iName) {
		_Router, _error := model.GetRoutersByNameAndNumberParts(offset, limit, iName,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByNameAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPathAndPatternOutlineHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iPatternOutline := self.Args("pattern_outline").String()

	if helper.IsHas(iPath) {
		_Router, _error := model.GetRoutersByPathAndPatternOutline(offset, limit, iPath,iPatternOutline)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPathAndPatternOutline's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPathAndFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iFit := self.Args("fit").MustInt()

	if helper.IsHas(iPath) {
		_Router, _error := model.GetRoutersByPathAndFit(offset, limit, iPath,iFit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPathAndFit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPathAndRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRoute := self.Args("route").Bytes()

	if helper.IsHas(iPath) {
		_Router, _error := model.GetRoutersByPathAndRoute(offset, limit, iPath,iRoute)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPathAndRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPathAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iPath) {
		_Router, _error := model.GetRoutersByPathAndNumberParts(offset, limit, iPath,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPathAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPatternOutlineAndFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPatternOutline := self.Args("pattern_outline").String()
	iFit := self.Args("fit").MustInt()

	if helper.IsHas(iPatternOutline) {
		_Router, _error := model.GetRoutersByPatternOutlineAndFit(offset, limit, iPatternOutline,iFit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPatternOutlineAndFit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPatternOutlineAndRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPatternOutline := self.Args("pattern_outline").String()
	iRoute := self.Args("route").Bytes()

	if helper.IsHas(iPatternOutline) {
		_Router, _error := model.GetRoutersByPatternOutlineAndRoute(offset, limit, iPatternOutline,iRoute)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPatternOutlineAndRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByPatternOutlineAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPatternOutline := self.Args("pattern_outline").String()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iPatternOutline) {
		_Router, _error := model.GetRoutersByPatternOutlineAndNumberParts(offset, limit, iPatternOutline,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByPatternOutlineAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByFitAndRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFit := self.Args("fit").MustInt()
	iRoute := self.Args("route").Bytes()

	if helper.IsHas(iFit) {
		_Router, _error := model.GetRoutersByFitAndRoute(offset, limit, iFit,iRoute)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByFitAndRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByFitAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFit := self.Args("fit").MustInt()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iFit) {
		_Router, _error := model.GetRoutersByFitAndNumberParts(offset, limit, iFit,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByFitAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersByRouteAndNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRoute := self.Args("route").Bytes()
	iNumberParts := self.Args("number_parts").MustInt()

	if helper.IsHas(iRoute) {
		_Router, _error := model.GetRoutersByRouteAndNumberParts(offset, limit, iRoute,iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRoutersByRouteAndNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRoutersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Router, _error := model.GetRouters(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRouters' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasRouterViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Router := model.HasRouterViaName(iName)
		var m = map[string]interface{}{}
		m["router"] = _Router
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasRouterViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasRouterViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPath := self.Args("path").String()
	if helper.IsHas(iPath) {
		_Router := model.HasRouterViaPath(iPath)
		var m = map[string]interface{}{}
		m["router"] = _Router
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasRouterViaPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasRouterViaPatternOutlineHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPatternOutline := self.Args("pattern_outline").String()
	if helper.IsHas(iPatternOutline) {
		_Router := model.HasRouterViaPatternOutline(iPatternOutline)
		var m = map[string]interface{}{}
		m["router"] = _Router
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasRouterViaPatternOutline's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasRouterViaFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFit := self.Args("fit").MustInt()
	if helper.IsHas(iFit) {
		_Router := model.HasRouterViaFit(iFit)
		var m = map[string]interface{}{}
		m["router"] = _Router
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasRouterViaFit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasRouterViaRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRoute := self.Args("route").Bytes()
	if helper.IsHas(iRoute) {
		_Router := model.HasRouterViaRoute(iRoute)
		var m = map[string]interface{}{}
		m["router"] = _Router
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasRouterViaRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasRouterViaNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNumberParts := self.Args("number_parts").MustInt()
	if helper.IsHas(iNumberParts) {
		_Router := model.HasRouterViaNumberParts(iNumberParts)
		var m = map[string]interface{}{}
		m["router"] = _Router
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasRouterViaNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRouterViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Router, _error := model.GetRouterViaName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRouterViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRouterViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPath := self.Args("path").String()
	if helper.IsHas(iPath) {
		_Router, _error := model.GetRouterViaPath(iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRouterViaPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRouterViaPatternOutlineHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPatternOutline := self.Args("pattern_outline").String()
	if helper.IsHas(iPatternOutline) {
		_Router, _error := model.GetRouterViaPatternOutline(iPatternOutline)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRouterViaPatternOutline's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRouterViaFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFit := self.Args("fit").MustInt()
	if helper.IsHas(iFit) {
		_Router, _error := model.GetRouterViaFit(iFit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRouterViaFit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRouterViaRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRoute := self.Args("route").Bytes()
	if helper.IsHas(iRoute) {
		_Router, _error := model.GetRouterViaRoute(iRoute)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRouterViaRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRouterViaNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNumberParts := self.Args("number_parts").MustInt()
	if helper.IsHas(iNumberParts) {
		_Router, _error := model.GetRouterViaNumberParts(iNumberParts)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the GetRouterViaNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetRouterViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iRouter model.Router
		self.Bind(&iRouter)
		_Router, _error := model.SetRouterViaName(Name_, &iRouter)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the SetRouterViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetRouterViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	if helper.IsHas(Path_) {
		var iRouter model.Router
		self.Bind(&iRouter)
		_Router, _error := model.SetRouterViaPath(Path_, &iRouter)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the SetRouterViaPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetRouterViaPatternOutlineHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PatternOutline_ := self.Args("pattern_outline").String()
	if helper.IsHas(PatternOutline_) {
		var iRouter model.Router
		self.Bind(&iRouter)
		_Router, _error := model.SetRouterViaPatternOutline(PatternOutline_, &iRouter)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the SetRouterViaPatternOutline's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetRouterViaFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fit_ := self.Args("fit").MustInt()
	if helper.IsHas(Fit_) {
		var iRouter model.Router
		self.Bind(&iRouter)
		_Router, _error := model.SetRouterViaFit(Fit_, &iRouter)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the SetRouterViaFit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetRouterViaRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Route_ := self.Args("route").Bytes()
	if helper.IsHas(Route_) {
		var iRouter model.Router
		self.Bind(&iRouter)
		_Router, _error := model.SetRouterViaRoute(Route_, &iRouter)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the SetRouterViaRoute's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetRouterViaNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	NumberParts_ := self.Args("number_parts").MustInt()
	if helper.IsHas(NumberParts_) {
		var iRouter model.Router
		self.Bind(&iRouter)
		_Router, _error := model.SetRouterViaNumberParts(NumberParts_, &iRouter)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Router)
	}
	herr.Message = "Can't get to the SetRouterViaNumberParts's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddRouterHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	Path_ := self.Args("path").String()
	PatternOutline_ := self.Args("pattern_outline").String()
	Fit_ := self.Args("fit").MustInt()
	Route_ := self.Args("route").Bytes()
	NumberParts_ := self.Args("number_parts").MustInt()

	if helper.IsHas(Name_) {
		_error := model.AddRouter(Name_,Path_,PatternOutline_,Fit_,Route_,NumberParts_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddRouter's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostRouterHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iRouter model.Router
	self.Bind(&iRouter)
	_string, _error := model.PostRouter(&iRouter)
	if (helper.IsHas(_string)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["created"] = _string
		return self.JSON(m, macross.StatusCreated)
	}
	return self.JSON(herr)
}

func PutRouterHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iRouter model.Router
	self.Bind(&iRouter)
	_string, _error := model.PutRouter(&iRouter)
	if (helper.IsHas(_string)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["updated"] = _string
		return self.JSON(m)
	}
	return self.JSON(herr)
}

func PutRouterViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iRouter model.Router
	self.Bind(&iRouter)
	_int64, _error := model.PutRouterViaName(Name_, &iRouter)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutRouterViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	var iRouter model.Router
	self.Bind(&iRouter)
	_int64, _error := model.PutRouterViaPath(Path_, &iRouter)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutRouterViaPatternOutlineHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PatternOutline_ := self.Args("pattern_outline").String()
	var iRouter model.Router
	self.Bind(&iRouter)
	_int64, _error := model.PutRouterViaPatternOutline(PatternOutline_, &iRouter)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutRouterViaFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fit_ := self.Args("fit").MustInt()
	var iRouter model.Router
	self.Bind(&iRouter)
	_int64, _error := model.PutRouterViaFit(Fit_, &iRouter)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutRouterViaRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Route_ := self.Args("route").Bytes()
	var iRouter model.Router
	self.Bind(&iRouter)
	_int64, _error := model.PutRouterViaRoute(Route_, &iRouter)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutRouterViaNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	NumberParts_ := self.Args("number_parts").MustInt()
	var iRouter model.Router
	self.Bind(&iRouter)
	_int64, _error := model.PutRouterViaNumberParts(NumberParts_, &iRouter)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateRouterViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iRouter model.Router
	self.Bind(&iRouter)
	var iMap = helper.StructToMap(iRouter)
	_error := model.UpdateRouterViaName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateRouterViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	var iRouter model.Router
	self.Bind(&iRouter)
	var iMap = helper.StructToMap(iRouter)
	_error := model.UpdateRouterViaPath(Path_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateRouterViaPatternOutlineHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PatternOutline_ := self.Args("pattern_outline").String()
	var iRouter model.Router
	self.Bind(&iRouter)
	var iMap = helper.StructToMap(iRouter)
	_error := model.UpdateRouterViaPatternOutline(PatternOutline_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateRouterViaFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fit_ := self.Args("fit").MustInt()
	var iRouter model.Router
	self.Bind(&iRouter)
	var iMap = helper.StructToMap(iRouter)
	_error := model.UpdateRouterViaFit(Fit_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateRouterViaRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Route_ := self.Args("route").Bytes()
	var iRouter model.Router
	self.Bind(&iRouter)
	var iMap = helper.StructToMap(iRouter)
	_error := model.UpdateRouterViaRoute(Route_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateRouterViaNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	NumberParts_ := self.Args("number_parts").MustInt()
	var iRouter model.Router
	self.Bind(&iRouter)
	var iMap = helper.StructToMap(iRouter)
	_error := model.UpdateRouterViaNumberParts(NumberParts_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteRouterHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_int64, _error := model.DeleteRouter(Name_)
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

func DeleteRouterViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteRouterViaName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteRouterViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	_error := model.DeleteRouterViaPath(Path_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteRouterViaPatternOutlineHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PatternOutline_ := self.Args("pattern_outline").String()
	_error := model.DeleteRouterViaPatternOutline(PatternOutline_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteRouterViaFitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fit_ := self.Args("fit").MustInt()
	_error := model.DeleteRouterViaFit(Fit_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteRouterViaRouteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Route_ := self.Args("route").Bytes()
	_error := model.DeleteRouterViaRoute(Route_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteRouterViaNumberPartsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	NumberParts_ := self.Args("number_parts").MustInt()
	_error := model.DeleteRouterViaNumberParts(NumberParts_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
