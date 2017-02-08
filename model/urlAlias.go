package model

type UrlAlias struct {
	Pid      int64  `xorm:"not null pk autoincr index(alias_langcode_pid) index(source_langcode_pid) INT(10)"`
	Source   string `xorm:"not null default '' index(source_langcode_pid) VARCHAR(255)"`
	Alias    string `xorm:"not null default '' index(alias_langcode_pid) VARCHAR(255)"`
	Langcode string `xorm:"not null default '' index(source_langcode_pid) index(alias_langcode_pid) VARCHAR(12)"`
}

// GetUrlAliasesCount UrlAliass' Count
func GetUrlAliasesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&UrlAlias{})
	return total, err
}

// GetUrlAliasCountViaPid Get UrlAlias via Pid
func GetUrlAliasCountViaPid(iPid int64) int64 {
	n, _ := Engine.Where("pid = ?", iPid).Count(&UrlAlias{Pid: iPid})
	return n
}

// GetUrlAliasCountViaSource Get UrlAlias via Source
func GetUrlAliasCountViaSource(iSource string) int64 {
	n, _ := Engine.Where("source = ?", iSource).Count(&UrlAlias{Source: iSource})
	return n
}

// GetUrlAliasCountViaAlias Get UrlAlias via Alias
func GetUrlAliasCountViaAlias(iAlias string) int64 {
	n, _ := Engine.Where("alias = ?", iAlias).Count(&UrlAlias{Alias: iAlias})
	return n
}

// GetUrlAliasCountViaLangcode Get UrlAlias via Langcode
func GetUrlAliasCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&UrlAlias{Langcode: iLangcode})
	return n
}

// GetUrlAliasesByPidAndSourceAndAlias Get UrlAliases via PidAndSourceAndAlias
func GetUrlAliasesByPidAndSourceAndAlias(offset int, limit int, Pid_ int64, Source_ string, Alias_ string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("pid = ? and source = ? and alias = ?", Pid_, Source_, Alias_).Limit(limit, offset).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesByPidAndSourceAndLangcode Get UrlAliases via PidAndSourceAndLangcode
func GetUrlAliasesByPidAndSourceAndLangcode(offset int, limit int, Pid_ int64, Source_ string, Langcode_ string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("pid = ? and source = ? and langcode = ?", Pid_, Source_, Langcode_).Limit(limit, offset).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesByPidAndAliasAndLangcode Get UrlAliases via PidAndAliasAndLangcode
func GetUrlAliasesByPidAndAliasAndLangcode(offset int, limit int, Pid_ int64, Alias_ string, Langcode_ string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("pid = ? and alias = ? and langcode = ?", Pid_, Alias_, Langcode_).Limit(limit, offset).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesBySourceAndAliasAndLangcode Get UrlAliases via SourceAndAliasAndLangcode
func GetUrlAliasesBySourceAndAliasAndLangcode(offset int, limit int, Source_ string, Alias_ string, Langcode_ string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("source = ? and alias = ? and langcode = ?", Source_, Alias_, Langcode_).Limit(limit, offset).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesByPidAndSource Get UrlAliases via PidAndSource
func GetUrlAliasesByPidAndSource(offset int, limit int, Pid_ int64, Source_ string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("pid = ? and source = ?", Pid_, Source_).Limit(limit, offset).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesByPidAndAlias Get UrlAliases via PidAndAlias
func GetUrlAliasesByPidAndAlias(offset int, limit int, Pid_ int64, Alias_ string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("pid = ? and alias = ?", Pid_, Alias_).Limit(limit, offset).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesByPidAndLangcode Get UrlAliases via PidAndLangcode
func GetUrlAliasesByPidAndLangcode(offset int, limit int, Pid_ int64, Langcode_ string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("pid = ? and langcode = ?", Pid_, Langcode_).Limit(limit, offset).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesBySourceAndAlias Get UrlAliases via SourceAndAlias
func GetUrlAliasesBySourceAndAlias(offset int, limit int, Source_ string, Alias_ string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("source = ? and alias = ?", Source_, Alias_).Limit(limit, offset).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesBySourceAndLangcode Get UrlAliases via SourceAndLangcode
func GetUrlAliasesBySourceAndLangcode(offset int, limit int, Source_ string, Langcode_ string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("source = ? and langcode = ?", Source_, Langcode_).Limit(limit, offset).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesByAliasAndLangcode Get UrlAliases via AliasAndLangcode
func GetUrlAliasesByAliasAndLangcode(offset int, limit int, Alias_ string, Langcode_ string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("alias = ? and langcode = ?", Alias_, Langcode_).Limit(limit, offset).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliases Get UrlAliases via field
func GetUrlAliases(offset int, limit int, field string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Limit(limit, offset).Desc(field).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesViaPid Get UrlAliass via Pid
func GetUrlAliasesViaPid(offset int, limit int, Pid_ int64, field string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("pid = ?", Pid_).Limit(limit, offset).Desc(field).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesViaSource Get UrlAliass via Source
func GetUrlAliasesViaSource(offset int, limit int, Source_ string, field string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("source = ?", Source_).Limit(limit, offset).Desc(field).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesViaAlias Get UrlAliass via Alias
func GetUrlAliasesViaAlias(offset int, limit int, Alias_ string, field string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("alias = ?", Alias_).Limit(limit, offset).Desc(field).Find(_UrlAlias)
	return _UrlAlias, err
}

// GetUrlAliasesViaLangcode Get UrlAliass via Langcode
func GetUrlAliasesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*UrlAlias, error) {
	var _UrlAlias = new([]*UrlAlias)
	err := Engine.Table("url_alias").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_UrlAlias)
	return _UrlAlias, err
}

// HasUrlAliasViaPid Has UrlAlias via Pid
func HasUrlAliasViaPid(iPid int64) bool {
	if has, err := Engine.Where("pid = ?", iPid).Get(new(UrlAlias)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUrlAliasViaSource Has UrlAlias via Source
func HasUrlAliasViaSource(iSource string) bool {
	if has, err := Engine.Where("source = ?", iSource).Get(new(UrlAlias)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUrlAliasViaAlias Has UrlAlias via Alias
func HasUrlAliasViaAlias(iAlias string) bool {
	if has, err := Engine.Where("alias = ?", iAlias).Get(new(UrlAlias)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUrlAliasViaLangcode Has UrlAlias via Langcode
func HasUrlAliasViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(UrlAlias)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetUrlAliasViaPid Get UrlAlias via Pid
func GetUrlAliasViaPid(iPid int64) (*UrlAlias, error) {
	var _UrlAlias = &UrlAlias{Pid: iPid}
	has, err := Engine.Get(_UrlAlias)
	if has {
		return _UrlAlias, err
	} else {
		return nil, err
	}
}

// GetUrlAliasViaSource Get UrlAlias via Source
func GetUrlAliasViaSource(iSource string) (*UrlAlias, error) {
	var _UrlAlias = &UrlAlias{Source: iSource}
	has, err := Engine.Get(_UrlAlias)
	if has {
		return _UrlAlias, err
	} else {
		return nil, err
	}
}

// GetUrlAliasViaAlias Get UrlAlias via Alias
func GetUrlAliasViaAlias(iAlias string) (*UrlAlias, error) {
	var _UrlAlias = &UrlAlias{Alias: iAlias}
	has, err := Engine.Get(_UrlAlias)
	if has {
		return _UrlAlias, err
	} else {
		return nil, err
	}
}

// GetUrlAliasViaLangcode Get UrlAlias via Langcode
func GetUrlAliasViaLangcode(iLangcode string) (*UrlAlias, error) {
	var _UrlAlias = &UrlAlias{Langcode: iLangcode}
	has, err := Engine.Get(_UrlAlias)
	if has {
		return _UrlAlias, err
	} else {
		return nil, err
	}
}

// SetUrlAliasViaPid Set UrlAlias via Pid
func SetUrlAliasViaPid(iPid int64, url_alias *UrlAlias) (int64, error) {
	url_alias.Pid = iPid
	return Engine.Insert(url_alias)
}

// SetUrlAliasViaSource Set UrlAlias via Source
func SetUrlAliasViaSource(iSource string, url_alias *UrlAlias) (int64, error) {
	url_alias.Source = iSource
	return Engine.Insert(url_alias)
}

// SetUrlAliasViaAlias Set UrlAlias via Alias
func SetUrlAliasViaAlias(iAlias string, url_alias *UrlAlias) (int64, error) {
	url_alias.Alias = iAlias
	return Engine.Insert(url_alias)
}

// SetUrlAliasViaLangcode Set UrlAlias via Langcode
func SetUrlAliasViaLangcode(iLangcode string, url_alias *UrlAlias) (int64, error) {
	url_alias.Langcode = iLangcode
	return Engine.Insert(url_alias)
}

// AddUrlAlias Add UrlAlias via all columns
func AddUrlAlias(iPid int64, iSource string, iAlias string, iLangcode string) error {
	_UrlAlias := &UrlAlias{Pid: iPid, Source: iSource, Alias: iAlias, Langcode: iLangcode}
	if _, err := Engine.Insert(_UrlAlias); err != nil {
		return err
	} else {
		return nil
	}
}

// PostUrlAlias Post UrlAlias via iUrlAlias
func PostUrlAlias(iUrlAlias *UrlAlias) (int64, error) {
	_, err := Engine.Insert(iUrlAlias)
	return iUrlAlias.Pid, err
}

// PutUrlAlias Put UrlAlias
func PutUrlAlias(iUrlAlias *UrlAlias) (int64, error) {
	_, err := Engine.Id(iUrlAlias.Pid).Update(iUrlAlias)
	return iUrlAlias.Pid, err
}

// PutUrlAliasViaPid Put UrlAlias via Pid
func PutUrlAliasViaPid(Pid_ int64, iUrlAlias *UrlAlias) (int64, error) {
	row, err := Engine.Update(iUrlAlias, &UrlAlias{Pid: Pid_})
	return row, err
}

// PutUrlAliasViaSource Put UrlAlias via Source
func PutUrlAliasViaSource(Source_ string, iUrlAlias *UrlAlias) (int64, error) {
	row, err := Engine.Update(iUrlAlias, &UrlAlias{Source: Source_})
	return row, err
}

// PutUrlAliasViaAlias Put UrlAlias via Alias
func PutUrlAliasViaAlias(Alias_ string, iUrlAlias *UrlAlias) (int64, error) {
	row, err := Engine.Update(iUrlAlias, &UrlAlias{Alias: Alias_})
	return row, err
}

// PutUrlAliasViaLangcode Put UrlAlias via Langcode
func PutUrlAliasViaLangcode(Langcode_ string, iUrlAlias *UrlAlias) (int64, error) {
	row, err := Engine.Update(iUrlAlias, &UrlAlias{Langcode: Langcode_})
	return row, err
}

// UpdateUrlAliasViaPid via map[string]interface{}{}
func UpdateUrlAliasViaPid(iPid int64, iUrlAliasMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UrlAlias)).Where("pid = ?", iPid).Update(iUrlAliasMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUrlAliasViaSource via map[string]interface{}{}
func UpdateUrlAliasViaSource(iSource string, iUrlAliasMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UrlAlias)).Where("source = ?", iSource).Update(iUrlAliasMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUrlAliasViaAlias via map[string]interface{}{}
func UpdateUrlAliasViaAlias(iAlias string, iUrlAliasMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UrlAlias)).Where("alias = ?", iAlias).Update(iUrlAliasMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUrlAliasViaLangcode via map[string]interface{}{}
func UpdateUrlAliasViaLangcode(iLangcode string, iUrlAliasMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UrlAlias)).Where("langcode = ?", iLangcode).Update(iUrlAliasMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteUrlAlias Delete UrlAlias
func DeleteUrlAlias(iPid int64) (int64, error) {
	row, err := Engine.Id(iPid).Delete(new(UrlAlias))
	return row, err
}

// DeleteUrlAliasViaPid Delete UrlAlias via Pid
func DeleteUrlAliasViaPid(iPid int64) (err error) {
	var has bool
	var _UrlAlias = &UrlAlias{Pid: iPid}
	if has, err = Engine.Get(_UrlAlias); (has == true) && (err == nil) {
		if row, err := Engine.Where("pid = ?", iPid).Delete(new(UrlAlias)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUrlAliasViaSource Delete UrlAlias via Source
func DeleteUrlAliasViaSource(iSource string) (err error) {
	var has bool
	var _UrlAlias = &UrlAlias{Source: iSource}
	if has, err = Engine.Get(_UrlAlias); (has == true) && (err == nil) {
		if row, err := Engine.Where("source = ?", iSource).Delete(new(UrlAlias)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUrlAliasViaAlias Delete UrlAlias via Alias
func DeleteUrlAliasViaAlias(iAlias string) (err error) {
	var has bool
	var _UrlAlias = &UrlAlias{Alias: iAlias}
	if has, err = Engine.Get(_UrlAlias); (has == true) && (err == nil) {
		if row, err := Engine.Where("alias = ?", iAlias).Delete(new(UrlAlias)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUrlAliasViaLangcode Delete UrlAlias via Langcode
func DeleteUrlAliasViaLangcode(iLangcode string) (err error) {
	var has bool
	var _UrlAlias = &UrlAlias{Langcode: iLangcode}
	if has, err = Engine.Get(_UrlAlias); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(UrlAlias)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
