package model

type Router struct {
	Name           string `xorm:"not null pk default '' VARCHAR(255)"`
	Path           string `xorm:"not null default '' VARCHAR(255)"`
	PatternOutline string `xorm:"not null default '' index(pattern_outline_parts) VARCHAR(255)"`
	Fit            int    `xorm:"not null default 0 INT(11)"`
	Route          []byte `xorm:"LONGBLOB"`
	NumberParts    int    `xorm:"not null default 0 index(pattern_outline_parts) SMALLINT(6)"`
}

// GetRoutersCount Routers' Count
func GetRoutersCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Router{})
	return total, err
}

// GetRouterCountViaName Get Router via Name
func GetRouterCountViaName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&Router{Name: iName})
	return n
}

// GetRouterCountViaPath Get Router via Path
func GetRouterCountViaPath(iPath string) int64 {
	n, _ := Engine.Where("path = ?", iPath).Count(&Router{Path: iPath})
	return n
}

// GetRouterCountViaPatternOutline Get Router via PatternOutline
func GetRouterCountViaPatternOutline(iPatternOutline string) int64 {
	n, _ := Engine.Where("pattern_outline = ?", iPatternOutline).Count(&Router{PatternOutline: iPatternOutline})
	return n
}

// GetRouterCountViaFit Get Router via Fit
func GetRouterCountViaFit(iFit int) int64 {
	n, _ := Engine.Where("fit = ?", iFit).Count(&Router{Fit: iFit})
	return n
}

// GetRouterCountViaRoute Get Router via Route
func GetRouterCountViaRoute(iRoute []byte) int64 {
	n, _ := Engine.Where("route = ?", iRoute).Count(&Router{Route: iRoute})
	return n
}

// GetRouterCountViaNumberParts Get Router via NumberParts
func GetRouterCountViaNumberParts(iNumberParts int) int64 {
	n, _ := Engine.Where("number_parts = ?", iNumberParts).Count(&Router{NumberParts: iNumberParts})
	return n
}

// GetRoutersByNameAndPathAndPatternOutline Get Routers via NameAndPathAndPatternOutline
func GetRoutersByNameAndPathAndPatternOutline(offset int, limit int, Name_ string, Path_ string, PatternOutline_ string) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and path = ? and pattern_outline = ?", Name_, Path_, PatternOutline_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndPathAndFit Get Routers via NameAndPathAndFit
func GetRoutersByNameAndPathAndFit(offset int, limit int, Name_ string, Path_ string, Fit_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and path = ? and fit = ?", Name_, Path_, Fit_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndPathAndRoute Get Routers via NameAndPathAndRoute
func GetRoutersByNameAndPathAndRoute(offset int, limit int, Name_ string, Path_ string, Route_ []byte) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and path = ? and route = ?", Name_, Path_, Route_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndPathAndNumberParts Get Routers via NameAndPathAndNumberParts
func GetRoutersByNameAndPathAndNumberParts(offset int, limit int, Name_ string, Path_ string, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and path = ? and number_parts = ?", Name_, Path_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndPatternOutlineAndFit Get Routers via NameAndPatternOutlineAndFit
func GetRoutersByNameAndPatternOutlineAndFit(offset int, limit int, Name_ string, PatternOutline_ string, Fit_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and pattern_outline = ? and fit = ?", Name_, PatternOutline_, Fit_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndPatternOutlineAndRoute Get Routers via NameAndPatternOutlineAndRoute
func GetRoutersByNameAndPatternOutlineAndRoute(offset int, limit int, Name_ string, PatternOutline_ string, Route_ []byte) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and pattern_outline = ? and route = ?", Name_, PatternOutline_, Route_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndPatternOutlineAndNumberParts Get Routers via NameAndPatternOutlineAndNumberParts
func GetRoutersByNameAndPatternOutlineAndNumberParts(offset int, limit int, Name_ string, PatternOutline_ string, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and pattern_outline = ? and number_parts = ?", Name_, PatternOutline_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndFitAndRoute Get Routers via NameAndFitAndRoute
func GetRoutersByNameAndFitAndRoute(offset int, limit int, Name_ string, Fit_ int, Route_ []byte) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and fit = ? and route = ?", Name_, Fit_, Route_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndFitAndNumberParts Get Routers via NameAndFitAndNumberParts
func GetRoutersByNameAndFitAndNumberParts(offset int, limit int, Name_ string, Fit_ int, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and fit = ? and number_parts = ?", Name_, Fit_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndRouteAndNumberParts Get Routers via NameAndRouteAndNumberParts
func GetRoutersByNameAndRouteAndNumberParts(offset int, limit int, Name_ string, Route_ []byte, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and route = ? and number_parts = ?", Name_, Route_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPathAndPatternOutlineAndFit Get Routers via PathAndPatternOutlineAndFit
func GetRoutersByPathAndPatternOutlineAndFit(offset int, limit int, Path_ string, PatternOutline_ string, Fit_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("path = ? and pattern_outline = ? and fit = ?", Path_, PatternOutline_, Fit_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPathAndPatternOutlineAndRoute Get Routers via PathAndPatternOutlineAndRoute
func GetRoutersByPathAndPatternOutlineAndRoute(offset int, limit int, Path_ string, PatternOutline_ string, Route_ []byte) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("path = ? and pattern_outline = ? and route = ?", Path_, PatternOutline_, Route_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPathAndPatternOutlineAndNumberParts Get Routers via PathAndPatternOutlineAndNumberParts
func GetRoutersByPathAndPatternOutlineAndNumberParts(offset int, limit int, Path_ string, PatternOutline_ string, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("path = ? and pattern_outline = ? and number_parts = ?", Path_, PatternOutline_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPathAndFitAndRoute Get Routers via PathAndFitAndRoute
func GetRoutersByPathAndFitAndRoute(offset int, limit int, Path_ string, Fit_ int, Route_ []byte) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("path = ? and fit = ? and route = ?", Path_, Fit_, Route_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPathAndFitAndNumberParts Get Routers via PathAndFitAndNumberParts
func GetRoutersByPathAndFitAndNumberParts(offset int, limit int, Path_ string, Fit_ int, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("path = ? and fit = ? and number_parts = ?", Path_, Fit_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPathAndRouteAndNumberParts Get Routers via PathAndRouteAndNumberParts
func GetRoutersByPathAndRouteAndNumberParts(offset int, limit int, Path_ string, Route_ []byte, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("path = ? and route = ? and number_parts = ?", Path_, Route_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPatternOutlineAndFitAndRoute Get Routers via PatternOutlineAndFitAndRoute
func GetRoutersByPatternOutlineAndFitAndRoute(offset int, limit int, PatternOutline_ string, Fit_ int, Route_ []byte) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("pattern_outline = ? and fit = ? and route = ?", PatternOutline_, Fit_, Route_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPatternOutlineAndFitAndNumberParts Get Routers via PatternOutlineAndFitAndNumberParts
func GetRoutersByPatternOutlineAndFitAndNumberParts(offset int, limit int, PatternOutline_ string, Fit_ int, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("pattern_outline = ? and fit = ? and number_parts = ?", PatternOutline_, Fit_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPatternOutlineAndRouteAndNumberParts Get Routers via PatternOutlineAndRouteAndNumberParts
func GetRoutersByPatternOutlineAndRouteAndNumberParts(offset int, limit int, PatternOutline_ string, Route_ []byte, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("pattern_outline = ? and route = ? and number_parts = ?", PatternOutline_, Route_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByFitAndRouteAndNumberParts Get Routers via FitAndRouteAndNumberParts
func GetRoutersByFitAndRouteAndNumberParts(offset int, limit int, Fit_ int, Route_ []byte, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("fit = ? and route = ? and number_parts = ?", Fit_, Route_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndPath Get Routers via NameAndPath
func GetRoutersByNameAndPath(offset int, limit int, Name_ string, Path_ string) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and path = ?", Name_, Path_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndPatternOutline Get Routers via NameAndPatternOutline
func GetRoutersByNameAndPatternOutline(offset int, limit int, Name_ string, PatternOutline_ string) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and pattern_outline = ?", Name_, PatternOutline_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndFit Get Routers via NameAndFit
func GetRoutersByNameAndFit(offset int, limit int, Name_ string, Fit_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and fit = ?", Name_, Fit_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndRoute Get Routers via NameAndRoute
func GetRoutersByNameAndRoute(offset int, limit int, Name_ string, Route_ []byte) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and route = ?", Name_, Route_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByNameAndNumberParts Get Routers via NameAndNumberParts
func GetRoutersByNameAndNumberParts(offset int, limit int, Name_ string, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ? and number_parts = ?", Name_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPathAndPatternOutline Get Routers via PathAndPatternOutline
func GetRoutersByPathAndPatternOutline(offset int, limit int, Path_ string, PatternOutline_ string) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("path = ? and pattern_outline = ?", Path_, PatternOutline_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPathAndFit Get Routers via PathAndFit
func GetRoutersByPathAndFit(offset int, limit int, Path_ string, Fit_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("path = ? and fit = ?", Path_, Fit_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPathAndRoute Get Routers via PathAndRoute
func GetRoutersByPathAndRoute(offset int, limit int, Path_ string, Route_ []byte) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("path = ? and route = ?", Path_, Route_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPathAndNumberParts Get Routers via PathAndNumberParts
func GetRoutersByPathAndNumberParts(offset int, limit int, Path_ string, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("path = ? and number_parts = ?", Path_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPatternOutlineAndFit Get Routers via PatternOutlineAndFit
func GetRoutersByPatternOutlineAndFit(offset int, limit int, PatternOutline_ string, Fit_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("pattern_outline = ? and fit = ?", PatternOutline_, Fit_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPatternOutlineAndRoute Get Routers via PatternOutlineAndRoute
func GetRoutersByPatternOutlineAndRoute(offset int, limit int, PatternOutline_ string, Route_ []byte) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("pattern_outline = ? and route = ?", PatternOutline_, Route_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByPatternOutlineAndNumberParts Get Routers via PatternOutlineAndNumberParts
func GetRoutersByPatternOutlineAndNumberParts(offset int, limit int, PatternOutline_ string, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("pattern_outline = ? and number_parts = ?", PatternOutline_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByFitAndRoute Get Routers via FitAndRoute
func GetRoutersByFitAndRoute(offset int, limit int, Fit_ int, Route_ []byte) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("fit = ? and route = ?", Fit_, Route_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByFitAndNumberParts Get Routers via FitAndNumberParts
func GetRoutersByFitAndNumberParts(offset int, limit int, Fit_ int, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("fit = ? and number_parts = ?", Fit_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRoutersByRouteAndNumberParts Get Routers via RouteAndNumberParts
func GetRoutersByRouteAndNumberParts(offset int, limit int, Route_ []byte, NumberParts_ int) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("route = ? and number_parts = ?", Route_, NumberParts_).Limit(limit, offset).Find(_Router)
	return _Router, err
}

// GetRouters Get Routers via field
func GetRouters(offset int, limit int, field string) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Limit(limit, offset).Desc(field).Find(_Router)
	return _Router, err
}

// GetRoutersViaName Get Routers via Name
func GetRoutersViaName(offset int, limit int, Name_ string, field string) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_Router)
	return _Router, err
}

// GetRoutersViaPath Get Routers via Path
func GetRoutersViaPath(offset int, limit int, Path_ string, field string) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("path = ?", Path_).Limit(limit, offset).Desc(field).Find(_Router)
	return _Router, err
}

// GetRoutersViaPatternOutline Get Routers via PatternOutline
func GetRoutersViaPatternOutline(offset int, limit int, PatternOutline_ string, field string) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("pattern_outline = ?", PatternOutline_).Limit(limit, offset).Desc(field).Find(_Router)
	return _Router, err
}

// GetRoutersViaFit Get Routers via Fit
func GetRoutersViaFit(offset int, limit int, Fit_ int, field string) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("fit = ?", Fit_).Limit(limit, offset).Desc(field).Find(_Router)
	return _Router, err
}

// GetRoutersViaRoute Get Routers via Route
func GetRoutersViaRoute(offset int, limit int, Route_ []byte, field string) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("route = ?", Route_).Limit(limit, offset).Desc(field).Find(_Router)
	return _Router, err
}

// GetRoutersViaNumberParts Get Routers via NumberParts
func GetRoutersViaNumberParts(offset int, limit int, NumberParts_ int, field string) (*[]*Router, error) {
	var _Router = new([]*Router)
	err := Engine.Table("router").Where("number_parts = ?", NumberParts_).Limit(limit, offset).Desc(field).Find(_Router)
	return _Router, err
}

// HasRouterViaName Has Router via Name
func HasRouterViaName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(Router)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRouterViaPath Has Router via Path
func HasRouterViaPath(iPath string) bool {
	if has, err := Engine.Where("path = ?", iPath).Get(new(Router)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRouterViaPatternOutline Has Router via PatternOutline
func HasRouterViaPatternOutline(iPatternOutline string) bool {
	if has, err := Engine.Where("pattern_outline = ?", iPatternOutline).Get(new(Router)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRouterViaFit Has Router via Fit
func HasRouterViaFit(iFit int) bool {
	if has, err := Engine.Where("fit = ?", iFit).Get(new(Router)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRouterViaRoute Has Router via Route
func HasRouterViaRoute(iRoute []byte) bool {
	if has, err := Engine.Where("route = ?", iRoute).Get(new(Router)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRouterViaNumberParts Has Router via NumberParts
func HasRouterViaNumberParts(iNumberParts int) bool {
	if has, err := Engine.Where("number_parts = ?", iNumberParts).Get(new(Router)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetRouterViaName Get Router via Name
func GetRouterViaName(iName string) (*Router, error) {
	var _Router = &Router{Name: iName}
	has, err := Engine.Get(_Router)
	if has {
		return _Router, err
	} else {
		return nil, err
	}
}

// GetRouterViaPath Get Router via Path
func GetRouterViaPath(iPath string) (*Router, error) {
	var _Router = &Router{Path: iPath}
	has, err := Engine.Get(_Router)
	if has {
		return _Router, err
	} else {
		return nil, err
	}
}

// GetRouterViaPatternOutline Get Router via PatternOutline
func GetRouterViaPatternOutline(iPatternOutline string) (*Router, error) {
	var _Router = &Router{PatternOutline: iPatternOutline}
	has, err := Engine.Get(_Router)
	if has {
		return _Router, err
	} else {
		return nil, err
	}
}

// GetRouterViaFit Get Router via Fit
func GetRouterViaFit(iFit int) (*Router, error) {
	var _Router = &Router{Fit: iFit}
	has, err := Engine.Get(_Router)
	if has {
		return _Router, err
	} else {
		return nil, err
	}
}

// GetRouterViaRoute Get Router via Route
func GetRouterViaRoute(iRoute []byte) (*Router, error) {
	var _Router = &Router{Route: iRoute}
	has, err := Engine.Get(_Router)
	if has {
		return _Router, err
	} else {
		return nil, err
	}
}

// GetRouterViaNumberParts Get Router via NumberParts
func GetRouterViaNumberParts(iNumberParts int) (*Router, error) {
	var _Router = &Router{NumberParts: iNumberParts}
	has, err := Engine.Get(_Router)
	if has {
		return _Router, err
	} else {
		return nil, err
	}
}

// SetRouterViaName Set Router via Name
func SetRouterViaName(iName string, router *Router) (int64, error) {
	router.Name = iName
	return Engine.Insert(router)
}

// SetRouterViaPath Set Router via Path
func SetRouterViaPath(iPath string, router *Router) (int64, error) {
	router.Path = iPath
	return Engine.Insert(router)
}

// SetRouterViaPatternOutline Set Router via PatternOutline
func SetRouterViaPatternOutline(iPatternOutline string, router *Router) (int64, error) {
	router.PatternOutline = iPatternOutline
	return Engine.Insert(router)
}

// SetRouterViaFit Set Router via Fit
func SetRouterViaFit(iFit int, router *Router) (int64, error) {
	router.Fit = iFit
	return Engine.Insert(router)
}

// SetRouterViaRoute Set Router via Route
func SetRouterViaRoute(iRoute []byte, router *Router) (int64, error) {
	router.Route = iRoute
	return Engine.Insert(router)
}

// SetRouterViaNumberParts Set Router via NumberParts
func SetRouterViaNumberParts(iNumberParts int, router *Router) (int64, error) {
	router.NumberParts = iNumberParts
	return Engine.Insert(router)
}

// AddRouter Add Router via all columns
func AddRouter(iName string, iPath string, iPatternOutline string, iFit int, iRoute []byte, iNumberParts int) error {
	_Router := &Router{Name: iName, Path: iPath, PatternOutline: iPatternOutline, Fit: iFit, Route: iRoute, NumberParts: iNumberParts}
	if _, err := Engine.Insert(_Router); err != nil {
		return err
	} else {
		return nil
	}
}

// PostRouter Post Router via iRouter
func PostRouter(iRouter *Router) (string, error) {
	_, err := Engine.Insert(iRouter)
	return iRouter.Name, err
}

// PutRouter Put Router
func PutRouter(iRouter *Router) (string, error) {
	_, err := Engine.Id(iRouter.Name).Update(iRouter)
	return iRouter.Name, err
}

// PutRouterViaName Put Router via Name
func PutRouterViaName(Name_ string, iRouter *Router) (int64, error) {
	row, err := Engine.Update(iRouter, &Router{Name: Name_})
	return row, err
}

// PutRouterViaPath Put Router via Path
func PutRouterViaPath(Path_ string, iRouter *Router) (int64, error) {
	row, err := Engine.Update(iRouter, &Router{Path: Path_})
	return row, err
}

// PutRouterViaPatternOutline Put Router via PatternOutline
func PutRouterViaPatternOutline(PatternOutline_ string, iRouter *Router) (int64, error) {
	row, err := Engine.Update(iRouter, &Router{PatternOutline: PatternOutline_})
	return row, err
}

// PutRouterViaFit Put Router via Fit
func PutRouterViaFit(Fit_ int, iRouter *Router) (int64, error) {
	row, err := Engine.Update(iRouter, &Router{Fit: Fit_})
	return row, err
}

// PutRouterViaRoute Put Router via Route
func PutRouterViaRoute(Route_ []byte, iRouter *Router) (int64, error) {
	row, err := Engine.Update(iRouter, &Router{Route: Route_})
	return row, err
}

// PutRouterViaNumberParts Put Router via NumberParts
func PutRouterViaNumberParts(NumberParts_ int, iRouter *Router) (int64, error) {
	row, err := Engine.Update(iRouter, &Router{NumberParts: NumberParts_})
	return row, err
}

// UpdateRouterViaName via map[string]interface{}{}
func UpdateRouterViaName(iName string, iRouterMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Router)).Where("name = ?", iName).Update(iRouterMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRouterViaPath via map[string]interface{}{}
func UpdateRouterViaPath(iPath string, iRouterMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Router)).Where("path = ?", iPath).Update(iRouterMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRouterViaPatternOutline via map[string]interface{}{}
func UpdateRouterViaPatternOutline(iPatternOutline string, iRouterMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Router)).Where("pattern_outline = ?", iPatternOutline).Update(iRouterMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRouterViaFit via map[string]interface{}{}
func UpdateRouterViaFit(iFit int, iRouterMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Router)).Where("fit = ?", iFit).Update(iRouterMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRouterViaRoute via map[string]interface{}{}
func UpdateRouterViaRoute(iRoute []byte, iRouterMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Router)).Where("route = ?", iRoute).Update(iRouterMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRouterViaNumberParts via map[string]interface{}{}
func UpdateRouterViaNumberParts(iNumberParts int, iRouterMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Router)).Where("number_parts = ?", iNumberParts).Update(iRouterMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteRouter Delete Router
func DeleteRouter(iName string) (int64, error) {
	row, err := Engine.Id(iName).Delete(new(Router))
	return row, err
}

// DeleteRouterViaName Delete Router via Name
func DeleteRouterViaName(iName string) (err error) {
	var has bool
	var _Router = &Router{Name: iName}
	if has, err = Engine.Get(_Router); (has == true) && (err == nil) {
		if row, err := Engine.Where("name = ?", iName).Delete(new(Router)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteRouterViaPath Delete Router via Path
func DeleteRouterViaPath(iPath string) (err error) {
	var has bool
	var _Router = &Router{Path: iPath}
	if has, err = Engine.Get(_Router); (has == true) && (err == nil) {
		if row, err := Engine.Where("path = ?", iPath).Delete(new(Router)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteRouterViaPatternOutline Delete Router via PatternOutline
func DeleteRouterViaPatternOutline(iPatternOutline string) (err error) {
	var has bool
	var _Router = &Router{PatternOutline: iPatternOutline}
	if has, err = Engine.Get(_Router); (has == true) && (err == nil) {
		if row, err := Engine.Where("pattern_outline = ?", iPatternOutline).Delete(new(Router)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteRouterViaFit Delete Router via Fit
func DeleteRouterViaFit(iFit int) (err error) {
	var has bool
	var _Router = &Router{Fit: iFit}
	if has, err = Engine.Get(_Router); (has == true) && (err == nil) {
		if row, err := Engine.Where("fit = ?", iFit).Delete(new(Router)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteRouterViaRoute Delete Router via Route
func DeleteRouterViaRoute(iRoute []byte) (err error) {
	var has bool
	var _Router = &Router{Route: iRoute}
	if has, err = Engine.Get(_Router); (has == true) && (err == nil) {
		if row, err := Engine.Where("route = ?", iRoute).Delete(new(Router)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteRouterViaNumberParts Delete Router via NumberParts
func DeleteRouterViaNumberParts(iNumberParts int) (err error) {
	var has bool
	var _Router = &Router{NumberParts: iNumberParts}
	if has, err = Engine.Get(_Router); (has == true) && (err == nil) {
		if row, err := Engine.Where("number_parts = ?", iNumberParts).Delete(new(Router)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
