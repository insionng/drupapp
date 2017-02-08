package model

type CacheConfig struct {
	Cid        string `xorm:"not null pk default '' VARCHAR(255)"`
	Data       []byte `xorm:"LONGBLOB"`
	Expire     int    `xorm:"not null default 0 index INT(11)"`
	Created    string `xorm:"not null default 0.000 DECIMAL(14,3)"`
	Serialized int    `xorm:"not null default 0 SMALLINT(6)"`
	Tags       string `xorm:"LONGTEXT"`
	Checksum   string `xorm:"not null VARCHAR(255)"`
}

// GetCacheConfigsCount CacheConfigs' Count
func GetCacheConfigsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CacheConfig{})
	return total, err
}

// GetCacheConfigCountViaCid Get CacheConfig via Cid
func GetCacheConfigCountViaCid(iCid string) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CacheConfig{Cid: iCid})
	return n
}

// GetCacheConfigCountViaData Get CacheConfig via Data
func GetCacheConfigCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&CacheConfig{Data: iData})
	return n
}

// GetCacheConfigCountViaExpire Get CacheConfig via Expire
func GetCacheConfigCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&CacheConfig{Expire: iExpire})
	return n
}

// GetCacheConfigCountViaCreated Get CacheConfig via Created
func GetCacheConfigCountViaCreated(iCreated string) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&CacheConfig{Created: iCreated})
	return n
}

// GetCacheConfigCountViaSerialized Get CacheConfig via Serialized
func GetCacheConfigCountViaSerialized(iSerialized int) int64 {
	n, _ := Engine.Where("serialized = ?", iSerialized).Count(&CacheConfig{Serialized: iSerialized})
	return n
}

// GetCacheConfigCountViaTags Get CacheConfig via Tags
func GetCacheConfigCountViaTags(iTags string) int64 {
	n, _ := Engine.Where("tags = ?", iTags).Count(&CacheConfig{Tags: iTags})
	return n
}

// GetCacheConfigCountViaChecksum Get CacheConfig via Checksum
func GetCacheConfigCountViaChecksum(iChecksum string) int64 {
	n, _ := Engine.Where("checksum = ?", iChecksum).Count(&CacheConfig{Checksum: iChecksum})
	return n
}

// GetCacheConfigsByCidAndDataAndExpire Get CacheConfigs via CidAndDataAndExpire
func GetCacheConfigsByCidAndDataAndExpire(offset int, limit int, Cid_ string, Data_ []byte, Expire_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and data = ? and expire = ?", Cid_, Data_, Expire_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndDataAndCreated Get CacheConfigs via CidAndDataAndCreated
func GetCacheConfigsByCidAndDataAndCreated(offset int, limit int, Cid_ string, Data_ []byte, Created_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and data = ? and created = ?", Cid_, Data_, Created_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndDataAndSerialized Get CacheConfigs via CidAndDataAndSerialized
func GetCacheConfigsByCidAndDataAndSerialized(offset int, limit int, Cid_ string, Data_ []byte, Serialized_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and data = ? and serialized = ?", Cid_, Data_, Serialized_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndDataAndTags Get CacheConfigs via CidAndDataAndTags
func GetCacheConfigsByCidAndDataAndTags(offset int, limit int, Cid_ string, Data_ []byte, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and data = ? and tags = ?", Cid_, Data_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndDataAndChecksum Get CacheConfigs via CidAndDataAndChecksum
func GetCacheConfigsByCidAndDataAndChecksum(offset int, limit int, Cid_ string, Data_ []byte, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and data = ? and checksum = ?", Cid_, Data_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndExpireAndCreated Get CacheConfigs via CidAndExpireAndCreated
func GetCacheConfigsByCidAndExpireAndCreated(offset int, limit int, Cid_ string, Expire_ int, Created_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and expire = ? and created = ?", Cid_, Expire_, Created_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndExpireAndSerialized Get CacheConfigs via CidAndExpireAndSerialized
func GetCacheConfigsByCidAndExpireAndSerialized(offset int, limit int, Cid_ string, Expire_ int, Serialized_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and expire = ? and serialized = ?", Cid_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndExpireAndTags Get CacheConfigs via CidAndExpireAndTags
func GetCacheConfigsByCidAndExpireAndTags(offset int, limit int, Cid_ string, Expire_ int, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and expire = ? and tags = ?", Cid_, Expire_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndExpireAndChecksum Get CacheConfigs via CidAndExpireAndChecksum
func GetCacheConfigsByCidAndExpireAndChecksum(offset int, limit int, Cid_ string, Expire_ int, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and expire = ? and checksum = ?", Cid_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndCreatedAndSerialized Get CacheConfigs via CidAndCreatedAndSerialized
func GetCacheConfigsByCidAndCreatedAndSerialized(offset int, limit int, Cid_ string, Created_ string, Serialized_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and created = ? and serialized = ?", Cid_, Created_, Serialized_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndCreatedAndTags Get CacheConfigs via CidAndCreatedAndTags
func GetCacheConfigsByCidAndCreatedAndTags(offset int, limit int, Cid_ string, Created_ string, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and created = ? and tags = ?", Cid_, Created_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndCreatedAndChecksum Get CacheConfigs via CidAndCreatedAndChecksum
func GetCacheConfigsByCidAndCreatedAndChecksum(offset int, limit int, Cid_ string, Created_ string, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and created = ? and checksum = ?", Cid_, Created_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndSerializedAndTags Get CacheConfigs via CidAndSerializedAndTags
func GetCacheConfigsByCidAndSerializedAndTags(offset int, limit int, Cid_ string, Serialized_ int, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and serialized = ? and tags = ?", Cid_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndSerializedAndChecksum Get CacheConfigs via CidAndSerializedAndChecksum
func GetCacheConfigsByCidAndSerializedAndChecksum(offset int, limit int, Cid_ string, Serialized_ int, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and serialized = ? and checksum = ?", Cid_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndTagsAndChecksum Get CacheConfigs via CidAndTagsAndChecksum
func GetCacheConfigsByCidAndTagsAndChecksum(offset int, limit int, Cid_ string, Tags_ string, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and tags = ? and checksum = ?", Cid_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndExpireAndCreated Get CacheConfigs via DataAndExpireAndCreated
func GetCacheConfigsByDataAndExpireAndCreated(offset int, limit int, Data_ []byte, Expire_ int, Created_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and expire = ? and created = ?", Data_, Expire_, Created_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndExpireAndSerialized Get CacheConfigs via DataAndExpireAndSerialized
func GetCacheConfigsByDataAndExpireAndSerialized(offset int, limit int, Data_ []byte, Expire_ int, Serialized_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and expire = ? and serialized = ?", Data_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndExpireAndTags Get CacheConfigs via DataAndExpireAndTags
func GetCacheConfigsByDataAndExpireAndTags(offset int, limit int, Data_ []byte, Expire_ int, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and expire = ? and tags = ?", Data_, Expire_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndExpireAndChecksum Get CacheConfigs via DataAndExpireAndChecksum
func GetCacheConfigsByDataAndExpireAndChecksum(offset int, limit int, Data_ []byte, Expire_ int, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and expire = ? and checksum = ?", Data_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndCreatedAndSerialized Get CacheConfigs via DataAndCreatedAndSerialized
func GetCacheConfigsByDataAndCreatedAndSerialized(offset int, limit int, Data_ []byte, Created_ string, Serialized_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and created = ? and serialized = ?", Data_, Created_, Serialized_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndCreatedAndTags Get CacheConfigs via DataAndCreatedAndTags
func GetCacheConfigsByDataAndCreatedAndTags(offset int, limit int, Data_ []byte, Created_ string, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and created = ? and tags = ?", Data_, Created_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndCreatedAndChecksum Get CacheConfigs via DataAndCreatedAndChecksum
func GetCacheConfigsByDataAndCreatedAndChecksum(offset int, limit int, Data_ []byte, Created_ string, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and created = ? and checksum = ?", Data_, Created_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndSerializedAndTags Get CacheConfigs via DataAndSerializedAndTags
func GetCacheConfigsByDataAndSerializedAndTags(offset int, limit int, Data_ []byte, Serialized_ int, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and serialized = ? and tags = ?", Data_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndSerializedAndChecksum Get CacheConfigs via DataAndSerializedAndChecksum
func GetCacheConfigsByDataAndSerializedAndChecksum(offset int, limit int, Data_ []byte, Serialized_ int, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and serialized = ? and checksum = ?", Data_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndTagsAndChecksum Get CacheConfigs via DataAndTagsAndChecksum
func GetCacheConfigsByDataAndTagsAndChecksum(offset int, limit int, Data_ []byte, Tags_ string, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and tags = ? and checksum = ?", Data_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByExpireAndCreatedAndSerialized Get CacheConfigs via ExpireAndCreatedAndSerialized
func GetCacheConfigsByExpireAndCreatedAndSerialized(offset int, limit int, Expire_ int, Created_ string, Serialized_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("expire = ? and created = ? and serialized = ?", Expire_, Created_, Serialized_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByExpireAndCreatedAndTags Get CacheConfigs via ExpireAndCreatedAndTags
func GetCacheConfigsByExpireAndCreatedAndTags(offset int, limit int, Expire_ int, Created_ string, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("expire = ? and created = ? and tags = ?", Expire_, Created_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByExpireAndCreatedAndChecksum Get CacheConfigs via ExpireAndCreatedAndChecksum
func GetCacheConfigsByExpireAndCreatedAndChecksum(offset int, limit int, Expire_ int, Created_ string, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("expire = ? and created = ? and checksum = ?", Expire_, Created_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByExpireAndSerializedAndTags Get CacheConfigs via ExpireAndSerializedAndTags
func GetCacheConfigsByExpireAndSerializedAndTags(offset int, limit int, Expire_ int, Serialized_ int, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("expire = ? and serialized = ? and tags = ?", Expire_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByExpireAndSerializedAndChecksum Get CacheConfigs via ExpireAndSerializedAndChecksum
func GetCacheConfigsByExpireAndSerializedAndChecksum(offset int, limit int, Expire_ int, Serialized_ int, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("expire = ? and serialized = ? and checksum = ?", Expire_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByExpireAndTagsAndChecksum Get CacheConfigs via ExpireAndTagsAndChecksum
func GetCacheConfigsByExpireAndTagsAndChecksum(offset int, limit int, Expire_ int, Tags_ string, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("expire = ? and tags = ? and checksum = ?", Expire_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCreatedAndSerializedAndTags Get CacheConfigs via CreatedAndSerializedAndTags
func GetCacheConfigsByCreatedAndSerializedAndTags(offset int, limit int, Created_ string, Serialized_ int, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("created = ? and serialized = ? and tags = ?", Created_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCreatedAndSerializedAndChecksum Get CacheConfigs via CreatedAndSerializedAndChecksum
func GetCacheConfigsByCreatedAndSerializedAndChecksum(offset int, limit int, Created_ string, Serialized_ int, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("created = ? and serialized = ? and checksum = ?", Created_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCreatedAndTagsAndChecksum Get CacheConfigs via CreatedAndTagsAndChecksum
func GetCacheConfigsByCreatedAndTagsAndChecksum(offset int, limit int, Created_ string, Tags_ string, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("created = ? and tags = ? and checksum = ?", Created_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsBySerializedAndTagsAndChecksum Get CacheConfigs via SerializedAndTagsAndChecksum
func GetCacheConfigsBySerializedAndTagsAndChecksum(offset int, limit int, Serialized_ int, Tags_ string, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("serialized = ? and tags = ? and checksum = ?", Serialized_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndData Get CacheConfigs via CidAndData
func GetCacheConfigsByCidAndData(offset int, limit int, Cid_ string, Data_ []byte) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and data = ?", Cid_, Data_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndExpire Get CacheConfigs via CidAndExpire
func GetCacheConfigsByCidAndExpire(offset int, limit int, Cid_ string, Expire_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and expire = ?", Cid_, Expire_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndCreated Get CacheConfigs via CidAndCreated
func GetCacheConfigsByCidAndCreated(offset int, limit int, Cid_ string, Created_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and created = ?", Cid_, Created_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndSerialized Get CacheConfigs via CidAndSerialized
func GetCacheConfigsByCidAndSerialized(offset int, limit int, Cid_ string, Serialized_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and serialized = ?", Cid_, Serialized_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndTags Get CacheConfigs via CidAndTags
func GetCacheConfigsByCidAndTags(offset int, limit int, Cid_ string, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and tags = ?", Cid_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCidAndChecksum Get CacheConfigs via CidAndChecksum
func GetCacheConfigsByCidAndChecksum(offset int, limit int, Cid_ string, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ? and checksum = ?", Cid_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndExpire Get CacheConfigs via DataAndExpire
func GetCacheConfigsByDataAndExpire(offset int, limit int, Data_ []byte, Expire_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and expire = ?", Data_, Expire_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndCreated Get CacheConfigs via DataAndCreated
func GetCacheConfigsByDataAndCreated(offset int, limit int, Data_ []byte, Created_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and created = ?", Data_, Created_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndSerialized Get CacheConfigs via DataAndSerialized
func GetCacheConfigsByDataAndSerialized(offset int, limit int, Data_ []byte, Serialized_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and serialized = ?", Data_, Serialized_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndTags Get CacheConfigs via DataAndTags
func GetCacheConfigsByDataAndTags(offset int, limit int, Data_ []byte, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and tags = ?", Data_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByDataAndChecksum Get CacheConfigs via DataAndChecksum
func GetCacheConfigsByDataAndChecksum(offset int, limit int, Data_ []byte, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ? and checksum = ?", Data_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByExpireAndCreated Get CacheConfigs via ExpireAndCreated
func GetCacheConfigsByExpireAndCreated(offset int, limit int, Expire_ int, Created_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("expire = ? and created = ?", Expire_, Created_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByExpireAndSerialized Get CacheConfigs via ExpireAndSerialized
func GetCacheConfigsByExpireAndSerialized(offset int, limit int, Expire_ int, Serialized_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("expire = ? and serialized = ?", Expire_, Serialized_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByExpireAndTags Get CacheConfigs via ExpireAndTags
func GetCacheConfigsByExpireAndTags(offset int, limit int, Expire_ int, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("expire = ? and tags = ?", Expire_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByExpireAndChecksum Get CacheConfigs via ExpireAndChecksum
func GetCacheConfigsByExpireAndChecksum(offset int, limit int, Expire_ int, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("expire = ? and checksum = ?", Expire_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCreatedAndSerialized Get CacheConfigs via CreatedAndSerialized
func GetCacheConfigsByCreatedAndSerialized(offset int, limit int, Created_ string, Serialized_ int) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("created = ? and serialized = ?", Created_, Serialized_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCreatedAndTags Get CacheConfigs via CreatedAndTags
func GetCacheConfigsByCreatedAndTags(offset int, limit int, Created_ string, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("created = ? and tags = ?", Created_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByCreatedAndChecksum Get CacheConfigs via CreatedAndChecksum
func GetCacheConfigsByCreatedAndChecksum(offset int, limit int, Created_ string, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("created = ? and checksum = ?", Created_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsBySerializedAndTags Get CacheConfigs via SerializedAndTags
func GetCacheConfigsBySerializedAndTags(offset int, limit int, Serialized_ int, Tags_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("serialized = ? and tags = ?", Serialized_, Tags_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsBySerializedAndChecksum Get CacheConfigs via SerializedAndChecksum
func GetCacheConfigsBySerializedAndChecksum(offset int, limit int, Serialized_ int, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("serialized = ? and checksum = ?", Serialized_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsByTagsAndChecksum Get CacheConfigs via TagsAndChecksum
func GetCacheConfigsByTagsAndChecksum(offset int, limit int, Tags_ string, Checksum_ string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("tags = ? and checksum = ?", Tags_, Checksum_).Limit(limit, offset).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigs Get CacheConfigs via field
func GetCacheConfigs(offset int, limit int, field string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Limit(limit, offset).Desc(field).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsViaCid Get CacheConfigs via Cid
func GetCacheConfigsViaCid(offset int, limit int, Cid_ string, field string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsViaData Get CacheConfigs via Data
func GetCacheConfigsViaData(offset int, limit int, Data_ []byte, field string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsViaExpire Get CacheConfigs via Expire
func GetCacheConfigsViaExpire(offset int, limit int, Expire_ int, field string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsViaCreated Get CacheConfigs via Created
func GetCacheConfigsViaCreated(offset int, limit int, Created_ string, field string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsViaSerialized Get CacheConfigs via Serialized
func GetCacheConfigsViaSerialized(offset int, limit int, Serialized_ int, field string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("serialized = ?", Serialized_).Limit(limit, offset).Desc(field).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsViaTags Get CacheConfigs via Tags
func GetCacheConfigsViaTags(offset int, limit int, Tags_ string, field string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("tags = ?", Tags_).Limit(limit, offset).Desc(field).Find(_CacheConfig)
	return _CacheConfig, err
}

// GetCacheConfigsViaChecksum Get CacheConfigs via Checksum
func GetCacheConfigsViaChecksum(offset int, limit int, Checksum_ string, field string) (*[]*CacheConfig, error) {
	var _CacheConfig = new([]*CacheConfig)
	err := Engine.Table("cache_config").Where("checksum = ?", Checksum_).Limit(limit, offset).Desc(field).Find(_CacheConfig)
	return _CacheConfig, err
}

// HasCacheConfigViaCid Has CacheConfig via Cid
func HasCacheConfigViaCid(iCid string) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CacheConfig)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheConfigViaData Has CacheConfig via Data
func HasCacheConfigViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(CacheConfig)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheConfigViaExpire Has CacheConfig via Expire
func HasCacheConfigViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(CacheConfig)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheConfigViaCreated Has CacheConfig via Created
func HasCacheConfigViaCreated(iCreated string) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(CacheConfig)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheConfigViaSerialized Has CacheConfig via Serialized
func HasCacheConfigViaSerialized(iSerialized int) bool {
	if has, err := Engine.Where("serialized = ?", iSerialized).Get(new(CacheConfig)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheConfigViaTags Has CacheConfig via Tags
func HasCacheConfigViaTags(iTags string) bool {
	if has, err := Engine.Where("tags = ?", iTags).Get(new(CacheConfig)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheConfigViaChecksum Has CacheConfig via Checksum
func HasCacheConfigViaChecksum(iChecksum string) bool {
	if has, err := Engine.Where("checksum = ?", iChecksum).Get(new(CacheConfig)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCacheConfigViaCid Get CacheConfig via Cid
func GetCacheConfigViaCid(iCid string) (*CacheConfig, error) {
	var _CacheConfig = &CacheConfig{Cid: iCid}
	has, err := Engine.Get(_CacheConfig)
	if has {
		return _CacheConfig, err
	} else {
		return nil, err
	}
}

// GetCacheConfigViaData Get CacheConfig via Data
func GetCacheConfigViaData(iData []byte) (*CacheConfig, error) {
	var _CacheConfig = &CacheConfig{Data: iData}
	has, err := Engine.Get(_CacheConfig)
	if has {
		return _CacheConfig, err
	} else {
		return nil, err
	}
}

// GetCacheConfigViaExpire Get CacheConfig via Expire
func GetCacheConfigViaExpire(iExpire int) (*CacheConfig, error) {
	var _CacheConfig = &CacheConfig{Expire: iExpire}
	has, err := Engine.Get(_CacheConfig)
	if has {
		return _CacheConfig, err
	} else {
		return nil, err
	}
}

// GetCacheConfigViaCreated Get CacheConfig via Created
func GetCacheConfigViaCreated(iCreated string) (*CacheConfig, error) {
	var _CacheConfig = &CacheConfig{Created: iCreated}
	has, err := Engine.Get(_CacheConfig)
	if has {
		return _CacheConfig, err
	} else {
		return nil, err
	}
}

// GetCacheConfigViaSerialized Get CacheConfig via Serialized
func GetCacheConfigViaSerialized(iSerialized int) (*CacheConfig, error) {
	var _CacheConfig = &CacheConfig{Serialized: iSerialized}
	has, err := Engine.Get(_CacheConfig)
	if has {
		return _CacheConfig, err
	} else {
		return nil, err
	}
}

// GetCacheConfigViaTags Get CacheConfig via Tags
func GetCacheConfigViaTags(iTags string) (*CacheConfig, error) {
	var _CacheConfig = &CacheConfig{Tags: iTags}
	has, err := Engine.Get(_CacheConfig)
	if has {
		return _CacheConfig, err
	} else {
		return nil, err
	}
}

// GetCacheConfigViaChecksum Get CacheConfig via Checksum
func GetCacheConfigViaChecksum(iChecksum string) (*CacheConfig, error) {
	var _CacheConfig = &CacheConfig{Checksum: iChecksum}
	has, err := Engine.Get(_CacheConfig)
	if has {
		return _CacheConfig, err
	} else {
		return nil, err
	}
}

// SetCacheConfigViaCid Set CacheConfig via Cid
func SetCacheConfigViaCid(iCid string, cache_config *CacheConfig) (int64, error) {
	cache_config.Cid = iCid
	return Engine.Insert(cache_config)
}

// SetCacheConfigViaData Set CacheConfig via Data
func SetCacheConfigViaData(iData []byte, cache_config *CacheConfig) (int64, error) {
	cache_config.Data = iData
	return Engine.Insert(cache_config)
}

// SetCacheConfigViaExpire Set CacheConfig via Expire
func SetCacheConfigViaExpire(iExpire int, cache_config *CacheConfig) (int64, error) {
	cache_config.Expire = iExpire
	return Engine.Insert(cache_config)
}

// SetCacheConfigViaCreated Set CacheConfig via Created
func SetCacheConfigViaCreated(iCreated string, cache_config *CacheConfig) (int64, error) {
	cache_config.Created = iCreated
	return Engine.Insert(cache_config)
}

// SetCacheConfigViaSerialized Set CacheConfig via Serialized
func SetCacheConfigViaSerialized(iSerialized int, cache_config *CacheConfig) (int64, error) {
	cache_config.Serialized = iSerialized
	return Engine.Insert(cache_config)
}

// SetCacheConfigViaTags Set CacheConfig via Tags
func SetCacheConfigViaTags(iTags string, cache_config *CacheConfig) (int64, error) {
	cache_config.Tags = iTags
	return Engine.Insert(cache_config)
}

// SetCacheConfigViaChecksum Set CacheConfig via Checksum
func SetCacheConfigViaChecksum(iChecksum string, cache_config *CacheConfig) (int64, error) {
	cache_config.Checksum = iChecksum
	return Engine.Insert(cache_config)
}

// AddCacheConfig Add CacheConfig via all columns
func AddCacheConfig(iCid string, iData []byte, iExpire int, iCreated string, iSerialized int, iTags string, iChecksum string) error {
	_CacheConfig := &CacheConfig{Cid: iCid, Data: iData, Expire: iExpire, Created: iCreated, Serialized: iSerialized, Tags: iTags, Checksum: iChecksum}
	if _, err := Engine.Insert(_CacheConfig); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCacheConfig Post CacheConfig via iCacheConfig
func PostCacheConfig(iCacheConfig *CacheConfig) (string, error) {
	_, err := Engine.Insert(iCacheConfig)
	return iCacheConfig.Cid, err
}

// PutCacheConfig Put CacheConfig
func PutCacheConfig(iCacheConfig *CacheConfig) (string, error) {
	_, err := Engine.Id(iCacheConfig.Cid).Update(iCacheConfig)
	return iCacheConfig.Cid, err
}

// PutCacheConfigViaCid Put CacheConfig via Cid
func PutCacheConfigViaCid(Cid_ string, iCacheConfig *CacheConfig) (int64, error) {
	row, err := Engine.Update(iCacheConfig, &CacheConfig{Cid: Cid_})
	return row, err
}

// PutCacheConfigViaData Put CacheConfig via Data
func PutCacheConfigViaData(Data_ []byte, iCacheConfig *CacheConfig) (int64, error) {
	row, err := Engine.Update(iCacheConfig, &CacheConfig{Data: Data_})
	return row, err
}

// PutCacheConfigViaExpire Put CacheConfig via Expire
func PutCacheConfigViaExpire(Expire_ int, iCacheConfig *CacheConfig) (int64, error) {
	row, err := Engine.Update(iCacheConfig, &CacheConfig{Expire: Expire_})
	return row, err
}

// PutCacheConfigViaCreated Put CacheConfig via Created
func PutCacheConfigViaCreated(Created_ string, iCacheConfig *CacheConfig) (int64, error) {
	row, err := Engine.Update(iCacheConfig, &CacheConfig{Created: Created_})
	return row, err
}

// PutCacheConfigViaSerialized Put CacheConfig via Serialized
func PutCacheConfigViaSerialized(Serialized_ int, iCacheConfig *CacheConfig) (int64, error) {
	row, err := Engine.Update(iCacheConfig, &CacheConfig{Serialized: Serialized_})
	return row, err
}

// PutCacheConfigViaTags Put CacheConfig via Tags
func PutCacheConfigViaTags(Tags_ string, iCacheConfig *CacheConfig) (int64, error) {
	row, err := Engine.Update(iCacheConfig, &CacheConfig{Tags: Tags_})
	return row, err
}

// PutCacheConfigViaChecksum Put CacheConfig via Checksum
func PutCacheConfigViaChecksum(Checksum_ string, iCacheConfig *CacheConfig) (int64, error) {
	row, err := Engine.Update(iCacheConfig, &CacheConfig{Checksum: Checksum_})
	return row, err
}

// UpdateCacheConfigViaCid via map[string]interface{}{}
func UpdateCacheConfigViaCid(iCid string, iCacheConfigMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheConfig)).Where("cid = ?", iCid).Update(iCacheConfigMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheConfigViaData via map[string]interface{}{}
func UpdateCacheConfigViaData(iData []byte, iCacheConfigMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheConfig)).Where("data = ?", iData).Update(iCacheConfigMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheConfigViaExpire via map[string]interface{}{}
func UpdateCacheConfigViaExpire(iExpire int, iCacheConfigMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheConfig)).Where("expire = ?", iExpire).Update(iCacheConfigMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheConfigViaCreated via map[string]interface{}{}
func UpdateCacheConfigViaCreated(iCreated string, iCacheConfigMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheConfig)).Where("created = ?", iCreated).Update(iCacheConfigMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheConfigViaSerialized via map[string]interface{}{}
func UpdateCacheConfigViaSerialized(iSerialized int, iCacheConfigMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheConfig)).Where("serialized = ?", iSerialized).Update(iCacheConfigMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheConfigViaTags via map[string]interface{}{}
func UpdateCacheConfigViaTags(iTags string, iCacheConfigMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheConfig)).Where("tags = ?", iTags).Update(iCacheConfigMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheConfigViaChecksum via map[string]interface{}{}
func UpdateCacheConfigViaChecksum(iChecksum string, iCacheConfigMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheConfig)).Where("checksum = ?", iChecksum).Update(iCacheConfigMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCacheConfig Delete CacheConfig
func DeleteCacheConfig(iCid string) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(CacheConfig))
	return row, err
}

// DeleteCacheConfigViaCid Delete CacheConfig via Cid
func DeleteCacheConfigViaCid(iCid string) (err error) {
	var has bool
	var _CacheConfig = &CacheConfig{Cid: iCid}
	if has, err = Engine.Get(_CacheConfig); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CacheConfig)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheConfigViaData Delete CacheConfig via Data
func DeleteCacheConfigViaData(iData []byte) (err error) {
	var has bool
	var _CacheConfig = &CacheConfig{Data: iData}
	if has, err = Engine.Get(_CacheConfig); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(CacheConfig)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheConfigViaExpire Delete CacheConfig via Expire
func DeleteCacheConfigViaExpire(iExpire int) (err error) {
	var has bool
	var _CacheConfig = &CacheConfig{Expire: iExpire}
	if has, err = Engine.Get(_CacheConfig); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(CacheConfig)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheConfigViaCreated Delete CacheConfig via Created
func DeleteCacheConfigViaCreated(iCreated string) (err error) {
	var has bool
	var _CacheConfig = &CacheConfig{Created: iCreated}
	if has, err = Engine.Get(_CacheConfig); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(CacheConfig)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheConfigViaSerialized Delete CacheConfig via Serialized
func DeleteCacheConfigViaSerialized(iSerialized int) (err error) {
	var has bool
	var _CacheConfig = &CacheConfig{Serialized: iSerialized}
	if has, err = Engine.Get(_CacheConfig); (has == true) && (err == nil) {
		if row, err := Engine.Where("serialized = ?", iSerialized).Delete(new(CacheConfig)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheConfigViaTags Delete CacheConfig via Tags
func DeleteCacheConfigViaTags(iTags string) (err error) {
	var has bool
	var _CacheConfig = &CacheConfig{Tags: iTags}
	if has, err = Engine.Get(_CacheConfig); (has == true) && (err == nil) {
		if row, err := Engine.Where("tags = ?", iTags).Delete(new(CacheConfig)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheConfigViaChecksum Delete CacheConfig via Checksum
func DeleteCacheConfigViaChecksum(iChecksum string) (err error) {
	var has bool
	var _CacheConfig = &CacheConfig{Checksum: iChecksum}
	if has, err = Engine.Get(_CacheConfig); (has == true) && (err == nil) {
		if row, err := Engine.Where("checksum = ?", iChecksum).Delete(new(CacheConfig)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
