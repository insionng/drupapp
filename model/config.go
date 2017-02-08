package model

type Config struct {
	Collection string `xorm:"not null pk default '' VARCHAR(255)"`
	Name       string `xorm:"not null pk default '' VARCHAR(255)"`
	Data       []byte `xorm:"LONGBLOB"`
}

// GetConfigsCount Configs' Count
func GetConfigsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Config{})
	return total, err
}

// GetConfigCountViaCollection Get Config via Collection
func GetConfigCountViaCollection(iCollection string) int64 {
	n, _ := Engine.Where("collection = ?", iCollection).Count(&Config{Collection: iCollection})
	return n
}

// GetConfigCountViaName Get Config via Name
func GetConfigCountViaName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&Config{Name: iName})
	return n
}

// GetConfigCountViaData Get Config via Data
func GetConfigCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&Config{Data: iData})
	return n
}

// GetConfigsByCollectionAndNameAndData Get Configs via CollectionAndNameAndData
func GetConfigsByCollectionAndNameAndData(offset int, limit int, Collection_ string, Name_ string, Data_ []byte) (*[]*Config, error) {
	var _Config = new([]*Config)
	err := Engine.Table("config").Where("collection = ? and name = ? and data = ?", Collection_, Name_, Data_).Limit(limit, offset).Find(_Config)
	return _Config, err
}

// GetConfigsByCollectionAndName Get Configs via CollectionAndName
func GetConfigsByCollectionAndName(offset int, limit int, Collection_ string, Name_ string) (*[]*Config, error) {
	var _Config = new([]*Config)
	err := Engine.Table("config").Where("collection = ? and name = ?", Collection_, Name_).Limit(limit, offset).Find(_Config)
	return _Config, err
}

// GetConfigsByCollectionAndData Get Configs via CollectionAndData
func GetConfigsByCollectionAndData(offset int, limit int, Collection_ string, Data_ []byte) (*[]*Config, error) {
	var _Config = new([]*Config)
	err := Engine.Table("config").Where("collection = ? and data = ?", Collection_, Data_).Limit(limit, offset).Find(_Config)
	return _Config, err
}

// GetConfigsByNameAndData Get Configs via NameAndData
func GetConfigsByNameAndData(offset int, limit int, Name_ string, Data_ []byte) (*[]*Config, error) {
	var _Config = new([]*Config)
	err := Engine.Table("config").Where("name = ? and data = ?", Name_, Data_).Limit(limit, offset).Find(_Config)
	return _Config, err
}

// GetConfigs Get Configs via field
func GetConfigs(offset int, limit int, field string) (*[]*Config, error) {
	var _Config = new([]*Config)
	err := Engine.Table("config").Limit(limit, offset).Desc(field).Find(_Config)
	return _Config, err
}

// GetConfigsViaCollection Get Configs via Collection
func GetConfigsViaCollection(offset int, limit int, Collection_ string, field string) (*[]*Config, error) {
	var _Config = new([]*Config)
	err := Engine.Table("config").Where("collection = ?", Collection_).Limit(limit, offset).Desc(field).Find(_Config)
	return _Config, err
}

// GetConfigsViaName Get Configs via Name
func GetConfigsViaName(offset int, limit int, Name_ string, field string) (*[]*Config, error) {
	var _Config = new([]*Config)
	err := Engine.Table("config").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_Config)
	return _Config, err
}

// GetConfigsViaData Get Configs via Data
func GetConfigsViaData(offset int, limit int, Data_ []byte, field string) (*[]*Config, error) {
	var _Config = new([]*Config)
	err := Engine.Table("config").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_Config)
	return _Config, err
}

// HasConfigViaCollection Has Config via Collection
func HasConfigViaCollection(iCollection string) bool {
	if has, err := Engine.Where("collection = ?", iCollection).Get(new(Config)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasConfigViaName Has Config via Name
func HasConfigViaName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(Config)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasConfigViaData Has Config via Data
func HasConfigViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(Config)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetConfigViaCollection Get Config via Collection
func GetConfigViaCollection(iCollection string) (*Config, error) {
	var _Config = &Config{Collection: iCollection}
	has, err := Engine.Get(_Config)
	if has {
		return _Config, err
	} else {
		return nil, err
	}
}

// GetConfigViaName Get Config via Name
func GetConfigViaName(iName string) (*Config, error) {
	var _Config = &Config{Name: iName}
	has, err := Engine.Get(_Config)
	if has {
		return _Config, err
	} else {
		return nil, err
	}
}

// GetConfigViaData Get Config via Data
func GetConfigViaData(iData []byte) (*Config, error) {
	var _Config = &Config{Data: iData}
	has, err := Engine.Get(_Config)
	if has {
		return _Config, err
	} else {
		return nil, err
	}
}

// SetConfigViaCollection Set Config via Collection
func SetConfigViaCollection(iCollection string, config *Config) (int64, error) {
	config.Collection = iCollection
	return Engine.Insert(config)
}

// SetConfigViaName Set Config via Name
func SetConfigViaName(iName string, config *Config) (int64, error) {
	config.Name = iName
	return Engine.Insert(config)
}

// SetConfigViaData Set Config via Data
func SetConfigViaData(iData []byte, config *Config) (int64, error) {
	config.Data = iData
	return Engine.Insert(config)
}

// AddConfig Add Config via all columns
func AddConfig(iCollection string, iName string, iData []byte) error {
	_Config := &Config{Collection: iCollection, Name: iName, Data: iData}
	if _, err := Engine.Insert(_Config); err != nil {
		return err
	} else {
		return nil
	}
}

// PostConfig Post Config via iConfig
func PostConfig(iConfig *Config) (string, error) {
	_, err := Engine.Insert(iConfig)
	return iConfig.Collection, err
}

// PutConfig Put Config
func PutConfig(iConfig *Config) (string, error) {
	_, err := Engine.Id(iConfig.Collection).Update(iConfig)
	return iConfig.Collection, err
}

// PutConfigViaCollection Put Config via Collection
func PutConfigViaCollection(Collection_ string, iConfig *Config) (int64, error) {
	row, err := Engine.Update(iConfig, &Config{Collection: Collection_})
	return row, err
}

// PutConfigViaName Put Config via Name
func PutConfigViaName(Name_ string, iConfig *Config) (int64, error) {
	row, err := Engine.Update(iConfig, &Config{Name: Name_})
	return row, err
}

// PutConfigViaData Put Config via Data
func PutConfigViaData(Data_ []byte, iConfig *Config) (int64, error) {
	row, err := Engine.Update(iConfig, &Config{Data: Data_})
	return row, err
}

// UpdateConfigViaCollection via map[string]interface{}{}
func UpdateConfigViaCollection(iCollection string, iConfigMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Config)).Where("collection = ?", iCollection).Update(iConfigMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateConfigViaName via map[string]interface{}{}
func UpdateConfigViaName(iName string, iConfigMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Config)).Where("name = ?", iName).Update(iConfigMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateConfigViaData via map[string]interface{}{}
func UpdateConfigViaData(iData []byte, iConfigMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Config)).Where("data = ?", iData).Update(iConfigMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteConfig Delete Config
func DeleteConfig(iCollection string) (int64, error) {
	row, err := Engine.Id(iCollection).Delete(new(Config))
	return row, err
}

// DeleteConfigViaCollection Delete Config via Collection
func DeleteConfigViaCollection(iCollection string) (err error) {
	var has bool
	var _Config = &Config{Collection: iCollection}
	if has, err = Engine.Get(_Config); (has == true) && (err == nil) {
		if row, err := Engine.Where("collection = ?", iCollection).Delete(new(Config)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteConfigViaName Delete Config via Name
func DeleteConfigViaName(iName string) (err error) {
	var has bool
	var _Config = &Config{Name: iName}
	if has, err = Engine.Get(_Config); (has == true) && (err == nil) {
		if row, err := Engine.Where("name = ?", iName).Delete(new(Config)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteConfigViaData Delete Config via Data
func DeleteConfigViaData(iData []byte) (err error) {
	var has bool
	var _Config = &Config{Data: iData}
	if has, err = Engine.Get(_Config); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(Config)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
