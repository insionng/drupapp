package model

type CacheDiscovery struct {
	Cid        string `xorm:"not null pk default '' VARCHAR(255)"`
	Data       []byte `xorm:"LONGBLOB"`
	Expire     int    `xorm:"not null default 0 index INT(11)"`
	Created    string `xorm:"not null default 0.000 DECIMAL(14,3)"`
	Serialized int    `xorm:"not null default 0 SMALLINT(6)"`
	Tags       string `xorm:"LONGTEXT"`
	Checksum   string `xorm:"not null VARCHAR(255)"`
}

// GetCacheDiscoveriesCount CacheDiscoverys' Count
func GetCacheDiscoveriesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CacheDiscovery{})
	return total, err
}

// GetCacheDiscoveryCountViaCid Get CacheDiscovery via Cid
func GetCacheDiscoveryCountViaCid(iCid string) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CacheDiscovery{Cid: iCid})
	return n
}

// GetCacheDiscoveryCountViaData Get CacheDiscovery via Data
func GetCacheDiscoveryCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&CacheDiscovery{Data: iData})
	return n
}

// GetCacheDiscoveryCountViaExpire Get CacheDiscovery via Expire
func GetCacheDiscoveryCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&CacheDiscovery{Expire: iExpire})
	return n
}

// GetCacheDiscoveryCountViaCreated Get CacheDiscovery via Created
func GetCacheDiscoveryCountViaCreated(iCreated string) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&CacheDiscovery{Created: iCreated})
	return n
}

// GetCacheDiscoveryCountViaSerialized Get CacheDiscovery via Serialized
func GetCacheDiscoveryCountViaSerialized(iSerialized int) int64 {
	n, _ := Engine.Where("serialized = ?", iSerialized).Count(&CacheDiscovery{Serialized: iSerialized})
	return n
}

// GetCacheDiscoveryCountViaTags Get CacheDiscovery via Tags
func GetCacheDiscoveryCountViaTags(iTags string) int64 {
	n, _ := Engine.Where("tags = ?", iTags).Count(&CacheDiscovery{Tags: iTags})
	return n
}

// GetCacheDiscoveryCountViaChecksum Get CacheDiscovery via Checksum
func GetCacheDiscoveryCountViaChecksum(iChecksum string) int64 {
	n, _ := Engine.Where("checksum = ?", iChecksum).Count(&CacheDiscovery{Checksum: iChecksum})
	return n
}

// GetCacheDiscoveriesByCidAndDataAndExpire Get CacheDiscoveries via CidAndDataAndExpire
func GetCacheDiscoveriesByCidAndDataAndExpire(offset int, limit int, Cid_ string, Data_ []byte, Expire_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and data = ? and expire = ?", Cid_, Data_, Expire_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndDataAndCreated Get CacheDiscoveries via CidAndDataAndCreated
func GetCacheDiscoveriesByCidAndDataAndCreated(offset int, limit int, Cid_ string, Data_ []byte, Created_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and data = ? and created = ?", Cid_, Data_, Created_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndDataAndSerialized Get CacheDiscoveries via CidAndDataAndSerialized
func GetCacheDiscoveriesByCidAndDataAndSerialized(offset int, limit int, Cid_ string, Data_ []byte, Serialized_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and data = ? and serialized = ?", Cid_, Data_, Serialized_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndDataAndTags Get CacheDiscoveries via CidAndDataAndTags
func GetCacheDiscoveriesByCidAndDataAndTags(offset int, limit int, Cid_ string, Data_ []byte, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and data = ? and tags = ?", Cid_, Data_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndDataAndChecksum Get CacheDiscoveries via CidAndDataAndChecksum
func GetCacheDiscoveriesByCidAndDataAndChecksum(offset int, limit int, Cid_ string, Data_ []byte, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and data = ? and checksum = ?", Cid_, Data_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndExpireAndCreated Get CacheDiscoveries via CidAndExpireAndCreated
func GetCacheDiscoveriesByCidAndExpireAndCreated(offset int, limit int, Cid_ string, Expire_ int, Created_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and expire = ? and created = ?", Cid_, Expire_, Created_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndExpireAndSerialized Get CacheDiscoveries via CidAndExpireAndSerialized
func GetCacheDiscoveriesByCidAndExpireAndSerialized(offset int, limit int, Cid_ string, Expire_ int, Serialized_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and expire = ? and serialized = ?", Cid_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndExpireAndTags Get CacheDiscoveries via CidAndExpireAndTags
func GetCacheDiscoveriesByCidAndExpireAndTags(offset int, limit int, Cid_ string, Expire_ int, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and expire = ? and tags = ?", Cid_, Expire_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndExpireAndChecksum Get CacheDiscoveries via CidAndExpireAndChecksum
func GetCacheDiscoveriesByCidAndExpireAndChecksum(offset int, limit int, Cid_ string, Expire_ int, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and expire = ? and checksum = ?", Cid_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndCreatedAndSerialized Get CacheDiscoveries via CidAndCreatedAndSerialized
func GetCacheDiscoveriesByCidAndCreatedAndSerialized(offset int, limit int, Cid_ string, Created_ string, Serialized_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and created = ? and serialized = ?", Cid_, Created_, Serialized_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndCreatedAndTags Get CacheDiscoveries via CidAndCreatedAndTags
func GetCacheDiscoveriesByCidAndCreatedAndTags(offset int, limit int, Cid_ string, Created_ string, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and created = ? and tags = ?", Cid_, Created_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndCreatedAndChecksum Get CacheDiscoveries via CidAndCreatedAndChecksum
func GetCacheDiscoveriesByCidAndCreatedAndChecksum(offset int, limit int, Cid_ string, Created_ string, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and created = ? and checksum = ?", Cid_, Created_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndSerializedAndTags Get CacheDiscoveries via CidAndSerializedAndTags
func GetCacheDiscoveriesByCidAndSerializedAndTags(offset int, limit int, Cid_ string, Serialized_ int, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and serialized = ? and tags = ?", Cid_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndSerializedAndChecksum Get CacheDiscoveries via CidAndSerializedAndChecksum
func GetCacheDiscoveriesByCidAndSerializedAndChecksum(offset int, limit int, Cid_ string, Serialized_ int, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and serialized = ? and checksum = ?", Cid_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndTagsAndChecksum Get CacheDiscoveries via CidAndTagsAndChecksum
func GetCacheDiscoveriesByCidAndTagsAndChecksum(offset int, limit int, Cid_ string, Tags_ string, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and tags = ? and checksum = ?", Cid_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndExpireAndCreated Get CacheDiscoveries via DataAndExpireAndCreated
func GetCacheDiscoveriesByDataAndExpireAndCreated(offset int, limit int, Data_ []byte, Expire_ int, Created_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and expire = ? and created = ?", Data_, Expire_, Created_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndExpireAndSerialized Get CacheDiscoveries via DataAndExpireAndSerialized
func GetCacheDiscoveriesByDataAndExpireAndSerialized(offset int, limit int, Data_ []byte, Expire_ int, Serialized_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and expire = ? and serialized = ?", Data_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndExpireAndTags Get CacheDiscoveries via DataAndExpireAndTags
func GetCacheDiscoveriesByDataAndExpireAndTags(offset int, limit int, Data_ []byte, Expire_ int, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and expire = ? and tags = ?", Data_, Expire_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndExpireAndChecksum Get CacheDiscoveries via DataAndExpireAndChecksum
func GetCacheDiscoveriesByDataAndExpireAndChecksum(offset int, limit int, Data_ []byte, Expire_ int, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and expire = ? and checksum = ?", Data_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndCreatedAndSerialized Get CacheDiscoveries via DataAndCreatedAndSerialized
func GetCacheDiscoveriesByDataAndCreatedAndSerialized(offset int, limit int, Data_ []byte, Created_ string, Serialized_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and created = ? and serialized = ?", Data_, Created_, Serialized_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndCreatedAndTags Get CacheDiscoveries via DataAndCreatedAndTags
func GetCacheDiscoveriesByDataAndCreatedAndTags(offset int, limit int, Data_ []byte, Created_ string, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and created = ? and tags = ?", Data_, Created_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndCreatedAndChecksum Get CacheDiscoveries via DataAndCreatedAndChecksum
func GetCacheDiscoveriesByDataAndCreatedAndChecksum(offset int, limit int, Data_ []byte, Created_ string, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and created = ? and checksum = ?", Data_, Created_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndSerializedAndTags Get CacheDiscoveries via DataAndSerializedAndTags
func GetCacheDiscoveriesByDataAndSerializedAndTags(offset int, limit int, Data_ []byte, Serialized_ int, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and serialized = ? and tags = ?", Data_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndSerializedAndChecksum Get CacheDiscoveries via DataAndSerializedAndChecksum
func GetCacheDiscoveriesByDataAndSerializedAndChecksum(offset int, limit int, Data_ []byte, Serialized_ int, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and serialized = ? and checksum = ?", Data_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndTagsAndChecksum Get CacheDiscoveries via DataAndTagsAndChecksum
func GetCacheDiscoveriesByDataAndTagsAndChecksum(offset int, limit int, Data_ []byte, Tags_ string, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and tags = ? and checksum = ?", Data_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByExpireAndCreatedAndSerialized Get CacheDiscoveries via ExpireAndCreatedAndSerialized
func GetCacheDiscoveriesByExpireAndCreatedAndSerialized(offset int, limit int, Expire_ int, Created_ string, Serialized_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("expire = ? and created = ? and serialized = ?", Expire_, Created_, Serialized_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByExpireAndCreatedAndTags Get CacheDiscoveries via ExpireAndCreatedAndTags
func GetCacheDiscoveriesByExpireAndCreatedAndTags(offset int, limit int, Expire_ int, Created_ string, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("expire = ? and created = ? and tags = ?", Expire_, Created_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByExpireAndCreatedAndChecksum Get CacheDiscoveries via ExpireAndCreatedAndChecksum
func GetCacheDiscoveriesByExpireAndCreatedAndChecksum(offset int, limit int, Expire_ int, Created_ string, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("expire = ? and created = ? and checksum = ?", Expire_, Created_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByExpireAndSerializedAndTags Get CacheDiscoveries via ExpireAndSerializedAndTags
func GetCacheDiscoveriesByExpireAndSerializedAndTags(offset int, limit int, Expire_ int, Serialized_ int, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("expire = ? and serialized = ? and tags = ?", Expire_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByExpireAndSerializedAndChecksum Get CacheDiscoveries via ExpireAndSerializedAndChecksum
func GetCacheDiscoveriesByExpireAndSerializedAndChecksum(offset int, limit int, Expire_ int, Serialized_ int, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("expire = ? and serialized = ? and checksum = ?", Expire_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByExpireAndTagsAndChecksum Get CacheDiscoveries via ExpireAndTagsAndChecksum
func GetCacheDiscoveriesByExpireAndTagsAndChecksum(offset int, limit int, Expire_ int, Tags_ string, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("expire = ? and tags = ? and checksum = ?", Expire_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCreatedAndSerializedAndTags Get CacheDiscoveries via CreatedAndSerializedAndTags
func GetCacheDiscoveriesByCreatedAndSerializedAndTags(offset int, limit int, Created_ string, Serialized_ int, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("created = ? and serialized = ? and tags = ?", Created_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCreatedAndSerializedAndChecksum Get CacheDiscoveries via CreatedAndSerializedAndChecksum
func GetCacheDiscoveriesByCreatedAndSerializedAndChecksum(offset int, limit int, Created_ string, Serialized_ int, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("created = ? and serialized = ? and checksum = ?", Created_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCreatedAndTagsAndChecksum Get CacheDiscoveries via CreatedAndTagsAndChecksum
func GetCacheDiscoveriesByCreatedAndTagsAndChecksum(offset int, limit int, Created_ string, Tags_ string, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("created = ? and tags = ? and checksum = ?", Created_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesBySerializedAndTagsAndChecksum Get CacheDiscoveries via SerializedAndTagsAndChecksum
func GetCacheDiscoveriesBySerializedAndTagsAndChecksum(offset int, limit int, Serialized_ int, Tags_ string, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("serialized = ? and tags = ? and checksum = ?", Serialized_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndData Get CacheDiscoveries via CidAndData
func GetCacheDiscoveriesByCidAndData(offset int, limit int, Cid_ string, Data_ []byte) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and data = ?", Cid_, Data_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndExpire Get CacheDiscoveries via CidAndExpire
func GetCacheDiscoveriesByCidAndExpire(offset int, limit int, Cid_ string, Expire_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and expire = ?", Cid_, Expire_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndCreated Get CacheDiscoveries via CidAndCreated
func GetCacheDiscoveriesByCidAndCreated(offset int, limit int, Cid_ string, Created_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and created = ?", Cid_, Created_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndSerialized Get CacheDiscoveries via CidAndSerialized
func GetCacheDiscoveriesByCidAndSerialized(offset int, limit int, Cid_ string, Serialized_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and serialized = ?", Cid_, Serialized_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndTags Get CacheDiscoveries via CidAndTags
func GetCacheDiscoveriesByCidAndTags(offset int, limit int, Cid_ string, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and tags = ?", Cid_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCidAndChecksum Get CacheDiscoveries via CidAndChecksum
func GetCacheDiscoveriesByCidAndChecksum(offset int, limit int, Cid_ string, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ? and checksum = ?", Cid_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndExpire Get CacheDiscoveries via DataAndExpire
func GetCacheDiscoveriesByDataAndExpire(offset int, limit int, Data_ []byte, Expire_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and expire = ?", Data_, Expire_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndCreated Get CacheDiscoveries via DataAndCreated
func GetCacheDiscoveriesByDataAndCreated(offset int, limit int, Data_ []byte, Created_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and created = ?", Data_, Created_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndSerialized Get CacheDiscoveries via DataAndSerialized
func GetCacheDiscoveriesByDataAndSerialized(offset int, limit int, Data_ []byte, Serialized_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and serialized = ?", Data_, Serialized_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndTags Get CacheDiscoveries via DataAndTags
func GetCacheDiscoveriesByDataAndTags(offset int, limit int, Data_ []byte, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and tags = ?", Data_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByDataAndChecksum Get CacheDiscoveries via DataAndChecksum
func GetCacheDiscoveriesByDataAndChecksum(offset int, limit int, Data_ []byte, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ? and checksum = ?", Data_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByExpireAndCreated Get CacheDiscoveries via ExpireAndCreated
func GetCacheDiscoveriesByExpireAndCreated(offset int, limit int, Expire_ int, Created_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("expire = ? and created = ?", Expire_, Created_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByExpireAndSerialized Get CacheDiscoveries via ExpireAndSerialized
func GetCacheDiscoveriesByExpireAndSerialized(offset int, limit int, Expire_ int, Serialized_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("expire = ? and serialized = ?", Expire_, Serialized_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByExpireAndTags Get CacheDiscoveries via ExpireAndTags
func GetCacheDiscoveriesByExpireAndTags(offset int, limit int, Expire_ int, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("expire = ? and tags = ?", Expire_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByExpireAndChecksum Get CacheDiscoveries via ExpireAndChecksum
func GetCacheDiscoveriesByExpireAndChecksum(offset int, limit int, Expire_ int, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("expire = ? and checksum = ?", Expire_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCreatedAndSerialized Get CacheDiscoveries via CreatedAndSerialized
func GetCacheDiscoveriesByCreatedAndSerialized(offset int, limit int, Created_ string, Serialized_ int) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("created = ? and serialized = ?", Created_, Serialized_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCreatedAndTags Get CacheDiscoveries via CreatedAndTags
func GetCacheDiscoveriesByCreatedAndTags(offset int, limit int, Created_ string, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("created = ? and tags = ?", Created_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByCreatedAndChecksum Get CacheDiscoveries via CreatedAndChecksum
func GetCacheDiscoveriesByCreatedAndChecksum(offset int, limit int, Created_ string, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("created = ? and checksum = ?", Created_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesBySerializedAndTags Get CacheDiscoveries via SerializedAndTags
func GetCacheDiscoveriesBySerializedAndTags(offset int, limit int, Serialized_ int, Tags_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("serialized = ? and tags = ?", Serialized_, Tags_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesBySerializedAndChecksum Get CacheDiscoveries via SerializedAndChecksum
func GetCacheDiscoveriesBySerializedAndChecksum(offset int, limit int, Serialized_ int, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("serialized = ? and checksum = ?", Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesByTagsAndChecksum Get CacheDiscoveries via TagsAndChecksum
func GetCacheDiscoveriesByTagsAndChecksum(offset int, limit int, Tags_ string, Checksum_ string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("tags = ? and checksum = ?", Tags_, Checksum_).Limit(limit, offset).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveries Get CacheDiscoveries via field
func GetCacheDiscoveries(offset int, limit int, field string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Limit(limit, offset).Desc(field).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesViaCid Get CacheDiscoverys via Cid
func GetCacheDiscoveriesViaCid(offset int, limit int, Cid_ string, field string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesViaData Get CacheDiscoverys via Data
func GetCacheDiscoveriesViaData(offset int, limit int, Data_ []byte, field string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesViaExpire Get CacheDiscoverys via Expire
func GetCacheDiscoveriesViaExpire(offset int, limit int, Expire_ int, field string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesViaCreated Get CacheDiscoverys via Created
func GetCacheDiscoveriesViaCreated(offset int, limit int, Created_ string, field string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesViaSerialized Get CacheDiscoverys via Serialized
func GetCacheDiscoveriesViaSerialized(offset int, limit int, Serialized_ int, field string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("serialized = ?", Serialized_).Limit(limit, offset).Desc(field).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesViaTags Get CacheDiscoverys via Tags
func GetCacheDiscoveriesViaTags(offset int, limit int, Tags_ string, field string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("tags = ?", Tags_).Limit(limit, offset).Desc(field).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// GetCacheDiscoveriesViaChecksum Get CacheDiscoverys via Checksum
func GetCacheDiscoveriesViaChecksum(offset int, limit int, Checksum_ string, field string) (*[]*CacheDiscovery, error) {
	var _CacheDiscovery = new([]*CacheDiscovery)
	err := Engine.Table("cache_discovery").Where("checksum = ?", Checksum_).Limit(limit, offset).Desc(field).Find(_CacheDiscovery)
	return _CacheDiscovery, err
}

// HasCacheDiscoveryViaCid Has CacheDiscovery via Cid
func HasCacheDiscoveryViaCid(iCid string) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CacheDiscovery)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDiscoveryViaData Has CacheDiscovery via Data
func HasCacheDiscoveryViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(CacheDiscovery)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDiscoveryViaExpire Has CacheDiscovery via Expire
func HasCacheDiscoveryViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(CacheDiscovery)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDiscoveryViaCreated Has CacheDiscovery via Created
func HasCacheDiscoveryViaCreated(iCreated string) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(CacheDiscovery)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDiscoveryViaSerialized Has CacheDiscovery via Serialized
func HasCacheDiscoveryViaSerialized(iSerialized int) bool {
	if has, err := Engine.Where("serialized = ?", iSerialized).Get(new(CacheDiscovery)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDiscoveryViaTags Has CacheDiscovery via Tags
func HasCacheDiscoveryViaTags(iTags string) bool {
	if has, err := Engine.Where("tags = ?", iTags).Get(new(CacheDiscovery)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDiscoveryViaChecksum Has CacheDiscovery via Checksum
func HasCacheDiscoveryViaChecksum(iChecksum string) bool {
	if has, err := Engine.Where("checksum = ?", iChecksum).Get(new(CacheDiscovery)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCacheDiscoveryViaCid Get CacheDiscovery via Cid
func GetCacheDiscoveryViaCid(iCid string) (*CacheDiscovery, error) {
	var _CacheDiscovery = &CacheDiscovery{Cid: iCid}
	has, err := Engine.Get(_CacheDiscovery)
	if has {
		return _CacheDiscovery, err
	} else {
		return nil, err
	}
}

// GetCacheDiscoveryViaData Get CacheDiscovery via Data
func GetCacheDiscoveryViaData(iData []byte) (*CacheDiscovery, error) {
	var _CacheDiscovery = &CacheDiscovery{Data: iData}
	has, err := Engine.Get(_CacheDiscovery)
	if has {
		return _CacheDiscovery, err
	} else {
		return nil, err
	}
}

// GetCacheDiscoveryViaExpire Get CacheDiscovery via Expire
func GetCacheDiscoveryViaExpire(iExpire int) (*CacheDiscovery, error) {
	var _CacheDiscovery = &CacheDiscovery{Expire: iExpire}
	has, err := Engine.Get(_CacheDiscovery)
	if has {
		return _CacheDiscovery, err
	} else {
		return nil, err
	}
}

// GetCacheDiscoveryViaCreated Get CacheDiscovery via Created
func GetCacheDiscoveryViaCreated(iCreated string) (*CacheDiscovery, error) {
	var _CacheDiscovery = &CacheDiscovery{Created: iCreated}
	has, err := Engine.Get(_CacheDiscovery)
	if has {
		return _CacheDiscovery, err
	} else {
		return nil, err
	}
}

// GetCacheDiscoveryViaSerialized Get CacheDiscovery via Serialized
func GetCacheDiscoveryViaSerialized(iSerialized int) (*CacheDiscovery, error) {
	var _CacheDiscovery = &CacheDiscovery{Serialized: iSerialized}
	has, err := Engine.Get(_CacheDiscovery)
	if has {
		return _CacheDiscovery, err
	} else {
		return nil, err
	}
}

// GetCacheDiscoveryViaTags Get CacheDiscovery via Tags
func GetCacheDiscoveryViaTags(iTags string) (*CacheDiscovery, error) {
	var _CacheDiscovery = &CacheDiscovery{Tags: iTags}
	has, err := Engine.Get(_CacheDiscovery)
	if has {
		return _CacheDiscovery, err
	} else {
		return nil, err
	}
}

// GetCacheDiscoveryViaChecksum Get CacheDiscovery via Checksum
func GetCacheDiscoveryViaChecksum(iChecksum string) (*CacheDiscovery, error) {
	var _CacheDiscovery = &CacheDiscovery{Checksum: iChecksum}
	has, err := Engine.Get(_CacheDiscovery)
	if has {
		return _CacheDiscovery, err
	} else {
		return nil, err
	}
}

// SetCacheDiscoveryViaCid Set CacheDiscovery via Cid
func SetCacheDiscoveryViaCid(iCid string, cache_discovery *CacheDiscovery) (int64, error) {
	cache_discovery.Cid = iCid
	return Engine.Insert(cache_discovery)
}

// SetCacheDiscoveryViaData Set CacheDiscovery via Data
func SetCacheDiscoveryViaData(iData []byte, cache_discovery *CacheDiscovery) (int64, error) {
	cache_discovery.Data = iData
	return Engine.Insert(cache_discovery)
}

// SetCacheDiscoveryViaExpire Set CacheDiscovery via Expire
func SetCacheDiscoveryViaExpire(iExpire int, cache_discovery *CacheDiscovery) (int64, error) {
	cache_discovery.Expire = iExpire
	return Engine.Insert(cache_discovery)
}

// SetCacheDiscoveryViaCreated Set CacheDiscovery via Created
func SetCacheDiscoveryViaCreated(iCreated string, cache_discovery *CacheDiscovery) (int64, error) {
	cache_discovery.Created = iCreated
	return Engine.Insert(cache_discovery)
}

// SetCacheDiscoveryViaSerialized Set CacheDiscovery via Serialized
func SetCacheDiscoveryViaSerialized(iSerialized int, cache_discovery *CacheDiscovery) (int64, error) {
	cache_discovery.Serialized = iSerialized
	return Engine.Insert(cache_discovery)
}

// SetCacheDiscoveryViaTags Set CacheDiscovery via Tags
func SetCacheDiscoveryViaTags(iTags string, cache_discovery *CacheDiscovery) (int64, error) {
	cache_discovery.Tags = iTags
	return Engine.Insert(cache_discovery)
}

// SetCacheDiscoveryViaChecksum Set CacheDiscovery via Checksum
func SetCacheDiscoveryViaChecksum(iChecksum string, cache_discovery *CacheDiscovery) (int64, error) {
	cache_discovery.Checksum = iChecksum
	return Engine.Insert(cache_discovery)
}

// AddCacheDiscovery Add CacheDiscovery via all columns
func AddCacheDiscovery(iCid string, iData []byte, iExpire int, iCreated string, iSerialized int, iTags string, iChecksum string) error {
	_CacheDiscovery := &CacheDiscovery{Cid: iCid, Data: iData, Expire: iExpire, Created: iCreated, Serialized: iSerialized, Tags: iTags, Checksum: iChecksum}
	if _, err := Engine.Insert(_CacheDiscovery); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCacheDiscovery Post CacheDiscovery via iCacheDiscovery
func PostCacheDiscovery(iCacheDiscovery *CacheDiscovery) (string, error) {
	_, err := Engine.Insert(iCacheDiscovery)
	return iCacheDiscovery.Cid, err
}

// PutCacheDiscovery Put CacheDiscovery
func PutCacheDiscovery(iCacheDiscovery *CacheDiscovery) (string, error) {
	_, err := Engine.Id(iCacheDiscovery.Cid).Update(iCacheDiscovery)
	return iCacheDiscovery.Cid, err
}

// PutCacheDiscoveryViaCid Put CacheDiscovery via Cid
func PutCacheDiscoveryViaCid(Cid_ string, iCacheDiscovery *CacheDiscovery) (int64, error) {
	row, err := Engine.Update(iCacheDiscovery, &CacheDiscovery{Cid: Cid_})
	return row, err
}

// PutCacheDiscoveryViaData Put CacheDiscovery via Data
func PutCacheDiscoveryViaData(Data_ []byte, iCacheDiscovery *CacheDiscovery) (int64, error) {
	row, err := Engine.Update(iCacheDiscovery, &CacheDiscovery{Data: Data_})
	return row, err
}

// PutCacheDiscoveryViaExpire Put CacheDiscovery via Expire
func PutCacheDiscoveryViaExpire(Expire_ int, iCacheDiscovery *CacheDiscovery) (int64, error) {
	row, err := Engine.Update(iCacheDiscovery, &CacheDiscovery{Expire: Expire_})
	return row, err
}

// PutCacheDiscoveryViaCreated Put CacheDiscovery via Created
func PutCacheDiscoveryViaCreated(Created_ string, iCacheDiscovery *CacheDiscovery) (int64, error) {
	row, err := Engine.Update(iCacheDiscovery, &CacheDiscovery{Created: Created_})
	return row, err
}

// PutCacheDiscoveryViaSerialized Put CacheDiscovery via Serialized
func PutCacheDiscoveryViaSerialized(Serialized_ int, iCacheDiscovery *CacheDiscovery) (int64, error) {
	row, err := Engine.Update(iCacheDiscovery, &CacheDiscovery{Serialized: Serialized_})
	return row, err
}

// PutCacheDiscoveryViaTags Put CacheDiscovery via Tags
func PutCacheDiscoveryViaTags(Tags_ string, iCacheDiscovery *CacheDiscovery) (int64, error) {
	row, err := Engine.Update(iCacheDiscovery, &CacheDiscovery{Tags: Tags_})
	return row, err
}

// PutCacheDiscoveryViaChecksum Put CacheDiscovery via Checksum
func PutCacheDiscoveryViaChecksum(Checksum_ string, iCacheDiscovery *CacheDiscovery) (int64, error) {
	row, err := Engine.Update(iCacheDiscovery, &CacheDiscovery{Checksum: Checksum_})
	return row, err
}

// UpdateCacheDiscoveryViaCid via map[string]interface{}{}
func UpdateCacheDiscoveryViaCid(iCid string, iCacheDiscoveryMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDiscovery)).Where("cid = ?", iCid).Update(iCacheDiscoveryMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDiscoveryViaData via map[string]interface{}{}
func UpdateCacheDiscoveryViaData(iData []byte, iCacheDiscoveryMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDiscovery)).Where("data = ?", iData).Update(iCacheDiscoveryMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDiscoveryViaExpire via map[string]interface{}{}
func UpdateCacheDiscoveryViaExpire(iExpire int, iCacheDiscoveryMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDiscovery)).Where("expire = ?", iExpire).Update(iCacheDiscoveryMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDiscoveryViaCreated via map[string]interface{}{}
func UpdateCacheDiscoveryViaCreated(iCreated string, iCacheDiscoveryMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDiscovery)).Where("created = ?", iCreated).Update(iCacheDiscoveryMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDiscoveryViaSerialized via map[string]interface{}{}
func UpdateCacheDiscoveryViaSerialized(iSerialized int, iCacheDiscoveryMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDiscovery)).Where("serialized = ?", iSerialized).Update(iCacheDiscoveryMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDiscoveryViaTags via map[string]interface{}{}
func UpdateCacheDiscoveryViaTags(iTags string, iCacheDiscoveryMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDiscovery)).Where("tags = ?", iTags).Update(iCacheDiscoveryMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDiscoveryViaChecksum via map[string]interface{}{}
func UpdateCacheDiscoveryViaChecksum(iChecksum string, iCacheDiscoveryMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDiscovery)).Where("checksum = ?", iChecksum).Update(iCacheDiscoveryMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCacheDiscovery Delete CacheDiscovery
func DeleteCacheDiscovery(iCid string) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(CacheDiscovery))
	return row, err
}

// DeleteCacheDiscoveryViaCid Delete CacheDiscovery via Cid
func DeleteCacheDiscoveryViaCid(iCid string) (err error) {
	var has bool
	var _CacheDiscovery = &CacheDiscovery{Cid: iCid}
	if has, err = Engine.Get(_CacheDiscovery); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CacheDiscovery)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDiscoveryViaData Delete CacheDiscovery via Data
func DeleteCacheDiscoveryViaData(iData []byte) (err error) {
	var has bool
	var _CacheDiscovery = &CacheDiscovery{Data: iData}
	if has, err = Engine.Get(_CacheDiscovery); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(CacheDiscovery)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDiscoveryViaExpire Delete CacheDiscovery via Expire
func DeleteCacheDiscoveryViaExpire(iExpire int) (err error) {
	var has bool
	var _CacheDiscovery = &CacheDiscovery{Expire: iExpire}
	if has, err = Engine.Get(_CacheDiscovery); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(CacheDiscovery)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDiscoveryViaCreated Delete CacheDiscovery via Created
func DeleteCacheDiscoveryViaCreated(iCreated string) (err error) {
	var has bool
	var _CacheDiscovery = &CacheDiscovery{Created: iCreated}
	if has, err = Engine.Get(_CacheDiscovery); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(CacheDiscovery)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDiscoveryViaSerialized Delete CacheDiscovery via Serialized
func DeleteCacheDiscoveryViaSerialized(iSerialized int) (err error) {
	var has bool
	var _CacheDiscovery = &CacheDiscovery{Serialized: iSerialized}
	if has, err = Engine.Get(_CacheDiscovery); (has == true) && (err == nil) {
		if row, err := Engine.Where("serialized = ?", iSerialized).Delete(new(CacheDiscovery)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDiscoveryViaTags Delete CacheDiscovery via Tags
func DeleteCacheDiscoveryViaTags(iTags string) (err error) {
	var has bool
	var _CacheDiscovery = &CacheDiscovery{Tags: iTags}
	if has, err = Engine.Get(_CacheDiscovery); (has == true) && (err == nil) {
		if row, err := Engine.Where("tags = ?", iTags).Delete(new(CacheDiscovery)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDiscoveryViaChecksum Delete CacheDiscovery via Checksum
func DeleteCacheDiscoveryViaChecksum(iChecksum string) (err error) {
	var has bool
	var _CacheDiscovery = &CacheDiscovery{Checksum: iChecksum}
	if has, err = Engine.Get(_CacheDiscovery); (has == true) && (err == nil) {
		if row, err := Engine.Where("checksum = ?", iChecksum).Delete(new(CacheDiscovery)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
