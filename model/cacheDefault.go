package model

type CacheDefault struct {
	Cid        string `xorm:"not null pk default '' VARCHAR(255)"`
	Data       []byte `xorm:"LONGBLOB"`
	Expire     int    `xorm:"not null default 0 index INT(11)"`
	Created    string `xorm:"not null default 0.000 DECIMAL(14,3)"`
	Serialized int    `xorm:"not null default 0 SMALLINT(6)"`
	Tags       string `xorm:"LONGTEXT"`
	Checksum   string `xorm:"not null VARCHAR(255)"`
}

// GetCacheDefaultsCount CacheDefaults' Count
func GetCacheDefaultsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CacheDefault{})
	return total, err
}

// GetCacheDefaultCountViaCid Get CacheDefault via Cid
func GetCacheDefaultCountViaCid(iCid string) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CacheDefault{Cid: iCid})
	return n
}

// GetCacheDefaultCountViaData Get CacheDefault via Data
func GetCacheDefaultCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&CacheDefault{Data: iData})
	return n
}

// GetCacheDefaultCountViaExpire Get CacheDefault via Expire
func GetCacheDefaultCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&CacheDefault{Expire: iExpire})
	return n
}

// GetCacheDefaultCountViaCreated Get CacheDefault via Created
func GetCacheDefaultCountViaCreated(iCreated string) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&CacheDefault{Created: iCreated})
	return n
}

// GetCacheDefaultCountViaSerialized Get CacheDefault via Serialized
func GetCacheDefaultCountViaSerialized(iSerialized int) int64 {
	n, _ := Engine.Where("serialized = ?", iSerialized).Count(&CacheDefault{Serialized: iSerialized})
	return n
}

// GetCacheDefaultCountViaTags Get CacheDefault via Tags
func GetCacheDefaultCountViaTags(iTags string) int64 {
	n, _ := Engine.Where("tags = ?", iTags).Count(&CacheDefault{Tags: iTags})
	return n
}

// GetCacheDefaultCountViaChecksum Get CacheDefault via Checksum
func GetCacheDefaultCountViaChecksum(iChecksum string) int64 {
	n, _ := Engine.Where("checksum = ?", iChecksum).Count(&CacheDefault{Checksum: iChecksum})
	return n
}

// GetCacheDefaultsByCidAndDataAndExpire Get CacheDefaults via CidAndDataAndExpire
func GetCacheDefaultsByCidAndDataAndExpire(offset int, limit int, Cid_ string, Data_ []byte, Expire_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and data = ? and expire = ?", Cid_, Data_, Expire_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndDataAndCreated Get CacheDefaults via CidAndDataAndCreated
func GetCacheDefaultsByCidAndDataAndCreated(offset int, limit int, Cid_ string, Data_ []byte, Created_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and data = ? and created = ?", Cid_, Data_, Created_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndDataAndSerialized Get CacheDefaults via CidAndDataAndSerialized
func GetCacheDefaultsByCidAndDataAndSerialized(offset int, limit int, Cid_ string, Data_ []byte, Serialized_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and data = ? and serialized = ?", Cid_, Data_, Serialized_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndDataAndTags Get CacheDefaults via CidAndDataAndTags
func GetCacheDefaultsByCidAndDataAndTags(offset int, limit int, Cid_ string, Data_ []byte, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and data = ? and tags = ?", Cid_, Data_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndDataAndChecksum Get CacheDefaults via CidAndDataAndChecksum
func GetCacheDefaultsByCidAndDataAndChecksum(offset int, limit int, Cid_ string, Data_ []byte, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and data = ? and checksum = ?", Cid_, Data_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndExpireAndCreated Get CacheDefaults via CidAndExpireAndCreated
func GetCacheDefaultsByCidAndExpireAndCreated(offset int, limit int, Cid_ string, Expire_ int, Created_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and expire = ? and created = ?", Cid_, Expire_, Created_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndExpireAndSerialized Get CacheDefaults via CidAndExpireAndSerialized
func GetCacheDefaultsByCidAndExpireAndSerialized(offset int, limit int, Cid_ string, Expire_ int, Serialized_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and expire = ? and serialized = ?", Cid_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndExpireAndTags Get CacheDefaults via CidAndExpireAndTags
func GetCacheDefaultsByCidAndExpireAndTags(offset int, limit int, Cid_ string, Expire_ int, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and expire = ? and tags = ?", Cid_, Expire_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndExpireAndChecksum Get CacheDefaults via CidAndExpireAndChecksum
func GetCacheDefaultsByCidAndExpireAndChecksum(offset int, limit int, Cid_ string, Expire_ int, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and expire = ? and checksum = ?", Cid_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndCreatedAndSerialized Get CacheDefaults via CidAndCreatedAndSerialized
func GetCacheDefaultsByCidAndCreatedAndSerialized(offset int, limit int, Cid_ string, Created_ string, Serialized_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and created = ? and serialized = ?", Cid_, Created_, Serialized_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndCreatedAndTags Get CacheDefaults via CidAndCreatedAndTags
func GetCacheDefaultsByCidAndCreatedAndTags(offset int, limit int, Cid_ string, Created_ string, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and created = ? and tags = ?", Cid_, Created_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndCreatedAndChecksum Get CacheDefaults via CidAndCreatedAndChecksum
func GetCacheDefaultsByCidAndCreatedAndChecksum(offset int, limit int, Cid_ string, Created_ string, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and created = ? and checksum = ?", Cid_, Created_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndSerializedAndTags Get CacheDefaults via CidAndSerializedAndTags
func GetCacheDefaultsByCidAndSerializedAndTags(offset int, limit int, Cid_ string, Serialized_ int, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and serialized = ? and tags = ?", Cid_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndSerializedAndChecksum Get CacheDefaults via CidAndSerializedAndChecksum
func GetCacheDefaultsByCidAndSerializedAndChecksum(offset int, limit int, Cid_ string, Serialized_ int, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and serialized = ? and checksum = ?", Cid_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndTagsAndChecksum Get CacheDefaults via CidAndTagsAndChecksum
func GetCacheDefaultsByCidAndTagsAndChecksum(offset int, limit int, Cid_ string, Tags_ string, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and tags = ? and checksum = ?", Cid_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndExpireAndCreated Get CacheDefaults via DataAndExpireAndCreated
func GetCacheDefaultsByDataAndExpireAndCreated(offset int, limit int, Data_ []byte, Expire_ int, Created_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and expire = ? and created = ?", Data_, Expire_, Created_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndExpireAndSerialized Get CacheDefaults via DataAndExpireAndSerialized
func GetCacheDefaultsByDataAndExpireAndSerialized(offset int, limit int, Data_ []byte, Expire_ int, Serialized_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and expire = ? and serialized = ?", Data_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndExpireAndTags Get CacheDefaults via DataAndExpireAndTags
func GetCacheDefaultsByDataAndExpireAndTags(offset int, limit int, Data_ []byte, Expire_ int, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and expire = ? and tags = ?", Data_, Expire_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndExpireAndChecksum Get CacheDefaults via DataAndExpireAndChecksum
func GetCacheDefaultsByDataAndExpireAndChecksum(offset int, limit int, Data_ []byte, Expire_ int, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and expire = ? and checksum = ?", Data_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndCreatedAndSerialized Get CacheDefaults via DataAndCreatedAndSerialized
func GetCacheDefaultsByDataAndCreatedAndSerialized(offset int, limit int, Data_ []byte, Created_ string, Serialized_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and created = ? and serialized = ?", Data_, Created_, Serialized_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndCreatedAndTags Get CacheDefaults via DataAndCreatedAndTags
func GetCacheDefaultsByDataAndCreatedAndTags(offset int, limit int, Data_ []byte, Created_ string, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and created = ? and tags = ?", Data_, Created_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndCreatedAndChecksum Get CacheDefaults via DataAndCreatedAndChecksum
func GetCacheDefaultsByDataAndCreatedAndChecksum(offset int, limit int, Data_ []byte, Created_ string, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and created = ? and checksum = ?", Data_, Created_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndSerializedAndTags Get CacheDefaults via DataAndSerializedAndTags
func GetCacheDefaultsByDataAndSerializedAndTags(offset int, limit int, Data_ []byte, Serialized_ int, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and serialized = ? and tags = ?", Data_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndSerializedAndChecksum Get CacheDefaults via DataAndSerializedAndChecksum
func GetCacheDefaultsByDataAndSerializedAndChecksum(offset int, limit int, Data_ []byte, Serialized_ int, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and serialized = ? and checksum = ?", Data_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndTagsAndChecksum Get CacheDefaults via DataAndTagsAndChecksum
func GetCacheDefaultsByDataAndTagsAndChecksum(offset int, limit int, Data_ []byte, Tags_ string, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and tags = ? and checksum = ?", Data_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByExpireAndCreatedAndSerialized Get CacheDefaults via ExpireAndCreatedAndSerialized
func GetCacheDefaultsByExpireAndCreatedAndSerialized(offset int, limit int, Expire_ int, Created_ string, Serialized_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("expire = ? and created = ? and serialized = ?", Expire_, Created_, Serialized_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByExpireAndCreatedAndTags Get CacheDefaults via ExpireAndCreatedAndTags
func GetCacheDefaultsByExpireAndCreatedAndTags(offset int, limit int, Expire_ int, Created_ string, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("expire = ? and created = ? and tags = ?", Expire_, Created_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByExpireAndCreatedAndChecksum Get CacheDefaults via ExpireAndCreatedAndChecksum
func GetCacheDefaultsByExpireAndCreatedAndChecksum(offset int, limit int, Expire_ int, Created_ string, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("expire = ? and created = ? and checksum = ?", Expire_, Created_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByExpireAndSerializedAndTags Get CacheDefaults via ExpireAndSerializedAndTags
func GetCacheDefaultsByExpireAndSerializedAndTags(offset int, limit int, Expire_ int, Serialized_ int, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("expire = ? and serialized = ? and tags = ?", Expire_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByExpireAndSerializedAndChecksum Get CacheDefaults via ExpireAndSerializedAndChecksum
func GetCacheDefaultsByExpireAndSerializedAndChecksum(offset int, limit int, Expire_ int, Serialized_ int, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("expire = ? and serialized = ? and checksum = ?", Expire_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByExpireAndTagsAndChecksum Get CacheDefaults via ExpireAndTagsAndChecksum
func GetCacheDefaultsByExpireAndTagsAndChecksum(offset int, limit int, Expire_ int, Tags_ string, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("expire = ? and tags = ? and checksum = ?", Expire_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCreatedAndSerializedAndTags Get CacheDefaults via CreatedAndSerializedAndTags
func GetCacheDefaultsByCreatedAndSerializedAndTags(offset int, limit int, Created_ string, Serialized_ int, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("created = ? and serialized = ? and tags = ?", Created_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCreatedAndSerializedAndChecksum Get CacheDefaults via CreatedAndSerializedAndChecksum
func GetCacheDefaultsByCreatedAndSerializedAndChecksum(offset int, limit int, Created_ string, Serialized_ int, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("created = ? and serialized = ? and checksum = ?", Created_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCreatedAndTagsAndChecksum Get CacheDefaults via CreatedAndTagsAndChecksum
func GetCacheDefaultsByCreatedAndTagsAndChecksum(offset int, limit int, Created_ string, Tags_ string, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("created = ? and tags = ? and checksum = ?", Created_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsBySerializedAndTagsAndChecksum Get CacheDefaults via SerializedAndTagsAndChecksum
func GetCacheDefaultsBySerializedAndTagsAndChecksum(offset int, limit int, Serialized_ int, Tags_ string, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("serialized = ? and tags = ? and checksum = ?", Serialized_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndData Get CacheDefaults via CidAndData
func GetCacheDefaultsByCidAndData(offset int, limit int, Cid_ string, Data_ []byte) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and data = ?", Cid_, Data_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndExpire Get CacheDefaults via CidAndExpire
func GetCacheDefaultsByCidAndExpire(offset int, limit int, Cid_ string, Expire_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and expire = ?", Cid_, Expire_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndCreated Get CacheDefaults via CidAndCreated
func GetCacheDefaultsByCidAndCreated(offset int, limit int, Cid_ string, Created_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and created = ?", Cid_, Created_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndSerialized Get CacheDefaults via CidAndSerialized
func GetCacheDefaultsByCidAndSerialized(offset int, limit int, Cid_ string, Serialized_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and serialized = ?", Cid_, Serialized_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndTags Get CacheDefaults via CidAndTags
func GetCacheDefaultsByCidAndTags(offset int, limit int, Cid_ string, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and tags = ?", Cid_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCidAndChecksum Get CacheDefaults via CidAndChecksum
func GetCacheDefaultsByCidAndChecksum(offset int, limit int, Cid_ string, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ? and checksum = ?", Cid_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndExpire Get CacheDefaults via DataAndExpire
func GetCacheDefaultsByDataAndExpire(offset int, limit int, Data_ []byte, Expire_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and expire = ?", Data_, Expire_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndCreated Get CacheDefaults via DataAndCreated
func GetCacheDefaultsByDataAndCreated(offset int, limit int, Data_ []byte, Created_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and created = ?", Data_, Created_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndSerialized Get CacheDefaults via DataAndSerialized
func GetCacheDefaultsByDataAndSerialized(offset int, limit int, Data_ []byte, Serialized_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and serialized = ?", Data_, Serialized_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndTags Get CacheDefaults via DataAndTags
func GetCacheDefaultsByDataAndTags(offset int, limit int, Data_ []byte, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and tags = ?", Data_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByDataAndChecksum Get CacheDefaults via DataAndChecksum
func GetCacheDefaultsByDataAndChecksum(offset int, limit int, Data_ []byte, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ? and checksum = ?", Data_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByExpireAndCreated Get CacheDefaults via ExpireAndCreated
func GetCacheDefaultsByExpireAndCreated(offset int, limit int, Expire_ int, Created_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("expire = ? and created = ?", Expire_, Created_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByExpireAndSerialized Get CacheDefaults via ExpireAndSerialized
func GetCacheDefaultsByExpireAndSerialized(offset int, limit int, Expire_ int, Serialized_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("expire = ? and serialized = ?", Expire_, Serialized_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByExpireAndTags Get CacheDefaults via ExpireAndTags
func GetCacheDefaultsByExpireAndTags(offset int, limit int, Expire_ int, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("expire = ? and tags = ?", Expire_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByExpireAndChecksum Get CacheDefaults via ExpireAndChecksum
func GetCacheDefaultsByExpireAndChecksum(offset int, limit int, Expire_ int, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("expire = ? and checksum = ?", Expire_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCreatedAndSerialized Get CacheDefaults via CreatedAndSerialized
func GetCacheDefaultsByCreatedAndSerialized(offset int, limit int, Created_ string, Serialized_ int) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("created = ? and serialized = ?", Created_, Serialized_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCreatedAndTags Get CacheDefaults via CreatedAndTags
func GetCacheDefaultsByCreatedAndTags(offset int, limit int, Created_ string, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("created = ? and tags = ?", Created_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByCreatedAndChecksum Get CacheDefaults via CreatedAndChecksum
func GetCacheDefaultsByCreatedAndChecksum(offset int, limit int, Created_ string, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("created = ? and checksum = ?", Created_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsBySerializedAndTags Get CacheDefaults via SerializedAndTags
func GetCacheDefaultsBySerializedAndTags(offset int, limit int, Serialized_ int, Tags_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("serialized = ? and tags = ?", Serialized_, Tags_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsBySerializedAndChecksum Get CacheDefaults via SerializedAndChecksum
func GetCacheDefaultsBySerializedAndChecksum(offset int, limit int, Serialized_ int, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("serialized = ? and checksum = ?", Serialized_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsByTagsAndChecksum Get CacheDefaults via TagsAndChecksum
func GetCacheDefaultsByTagsAndChecksum(offset int, limit int, Tags_ string, Checksum_ string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("tags = ? and checksum = ?", Tags_, Checksum_).Limit(limit, offset).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaults Get CacheDefaults via field
func GetCacheDefaults(offset int, limit int, field string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Limit(limit, offset).Desc(field).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsViaCid Get CacheDefaults via Cid
func GetCacheDefaultsViaCid(offset int, limit int, Cid_ string, field string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsViaData Get CacheDefaults via Data
func GetCacheDefaultsViaData(offset int, limit int, Data_ []byte, field string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsViaExpire Get CacheDefaults via Expire
func GetCacheDefaultsViaExpire(offset int, limit int, Expire_ int, field string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsViaCreated Get CacheDefaults via Created
func GetCacheDefaultsViaCreated(offset int, limit int, Created_ string, field string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsViaSerialized Get CacheDefaults via Serialized
func GetCacheDefaultsViaSerialized(offset int, limit int, Serialized_ int, field string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("serialized = ?", Serialized_).Limit(limit, offset).Desc(field).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsViaTags Get CacheDefaults via Tags
func GetCacheDefaultsViaTags(offset int, limit int, Tags_ string, field string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("tags = ?", Tags_).Limit(limit, offset).Desc(field).Find(_CacheDefault)
	return _CacheDefault, err
}

// GetCacheDefaultsViaChecksum Get CacheDefaults via Checksum
func GetCacheDefaultsViaChecksum(offset int, limit int, Checksum_ string, field string) (*[]*CacheDefault, error) {
	var _CacheDefault = new([]*CacheDefault)
	err := Engine.Table("cache_default").Where("checksum = ?", Checksum_).Limit(limit, offset).Desc(field).Find(_CacheDefault)
	return _CacheDefault, err
}

// HasCacheDefaultViaCid Has CacheDefault via Cid
func HasCacheDefaultViaCid(iCid string) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CacheDefault)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDefaultViaData Has CacheDefault via Data
func HasCacheDefaultViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(CacheDefault)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDefaultViaExpire Has CacheDefault via Expire
func HasCacheDefaultViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(CacheDefault)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDefaultViaCreated Has CacheDefault via Created
func HasCacheDefaultViaCreated(iCreated string) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(CacheDefault)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDefaultViaSerialized Has CacheDefault via Serialized
func HasCacheDefaultViaSerialized(iSerialized int) bool {
	if has, err := Engine.Where("serialized = ?", iSerialized).Get(new(CacheDefault)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDefaultViaTags Has CacheDefault via Tags
func HasCacheDefaultViaTags(iTags string) bool {
	if has, err := Engine.Where("tags = ?", iTags).Get(new(CacheDefault)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheDefaultViaChecksum Has CacheDefault via Checksum
func HasCacheDefaultViaChecksum(iChecksum string) bool {
	if has, err := Engine.Where("checksum = ?", iChecksum).Get(new(CacheDefault)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCacheDefaultViaCid Get CacheDefault via Cid
func GetCacheDefaultViaCid(iCid string) (*CacheDefault, error) {
	var _CacheDefault = &CacheDefault{Cid: iCid}
	has, err := Engine.Get(_CacheDefault)
	if has {
		return _CacheDefault, err
	} else {
		return nil, err
	}
}

// GetCacheDefaultViaData Get CacheDefault via Data
func GetCacheDefaultViaData(iData []byte) (*CacheDefault, error) {
	var _CacheDefault = &CacheDefault{Data: iData}
	has, err := Engine.Get(_CacheDefault)
	if has {
		return _CacheDefault, err
	} else {
		return nil, err
	}
}

// GetCacheDefaultViaExpire Get CacheDefault via Expire
func GetCacheDefaultViaExpire(iExpire int) (*CacheDefault, error) {
	var _CacheDefault = &CacheDefault{Expire: iExpire}
	has, err := Engine.Get(_CacheDefault)
	if has {
		return _CacheDefault, err
	} else {
		return nil, err
	}
}

// GetCacheDefaultViaCreated Get CacheDefault via Created
func GetCacheDefaultViaCreated(iCreated string) (*CacheDefault, error) {
	var _CacheDefault = &CacheDefault{Created: iCreated}
	has, err := Engine.Get(_CacheDefault)
	if has {
		return _CacheDefault, err
	} else {
		return nil, err
	}
}

// GetCacheDefaultViaSerialized Get CacheDefault via Serialized
func GetCacheDefaultViaSerialized(iSerialized int) (*CacheDefault, error) {
	var _CacheDefault = &CacheDefault{Serialized: iSerialized}
	has, err := Engine.Get(_CacheDefault)
	if has {
		return _CacheDefault, err
	} else {
		return nil, err
	}
}

// GetCacheDefaultViaTags Get CacheDefault via Tags
func GetCacheDefaultViaTags(iTags string) (*CacheDefault, error) {
	var _CacheDefault = &CacheDefault{Tags: iTags}
	has, err := Engine.Get(_CacheDefault)
	if has {
		return _CacheDefault, err
	} else {
		return nil, err
	}
}

// GetCacheDefaultViaChecksum Get CacheDefault via Checksum
func GetCacheDefaultViaChecksum(iChecksum string) (*CacheDefault, error) {
	var _CacheDefault = &CacheDefault{Checksum: iChecksum}
	has, err := Engine.Get(_CacheDefault)
	if has {
		return _CacheDefault, err
	} else {
		return nil, err
	}
}

// SetCacheDefaultViaCid Set CacheDefault via Cid
func SetCacheDefaultViaCid(iCid string, cache_default *CacheDefault) (int64, error) {
	cache_default.Cid = iCid
	return Engine.Insert(cache_default)
}

// SetCacheDefaultViaData Set CacheDefault via Data
func SetCacheDefaultViaData(iData []byte, cache_default *CacheDefault) (int64, error) {
	cache_default.Data = iData
	return Engine.Insert(cache_default)
}

// SetCacheDefaultViaExpire Set CacheDefault via Expire
func SetCacheDefaultViaExpire(iExpire int, cache_default *CacheDefault) (int64, error) {
	cache_default.Expire = iExpire
	return Engine.Insert(cache_default)
}

// SetCacheDefaultViaCreated Set CacheDefault via Created
func SetCacheDefaultViaCreated(iCreated string, cache_default *CacheDefault) (int64, error) {
	cache_default.Created = iCreated
	return Engine.Insert(cache_default)
}

// SetCacheDefaultViaSerialized Set CacheDefault via Serialized
func SetCacheDefaultViaSerialized(iSerialized int, cache_default *CacheDefault) (int64, error) {
	cache_default.Serialized = iSerialized
	return Engine.Insert(cache_default)
}

// SetCacheDefaultViaTags Set CacheDefault via Tags
func SetCacheDefaultViaTags(iTags string, cache_default *CacheDefault) (int64, error) {
	cache_default.Tags = iTags
	return Engine.Insert(cache_default)
}

// SetCacheDefaultViaChecksum Set CacheDefault via Checksum
func SetCacheDefaultViaChecksum(iChecksum string, cache_default *CacheDefault) (int64, error) {
	cache_default.Checksum = iChecksum
	return Engine.Insert(cache_default)
}

// AddCacheDefault Add CacheDefault via all columns
func AddCacheDefault(iCid string, iData []byte, iExpire int, iCreated string, iSerialized int, iTags string, iChecksum string) error {
	_CacheDefault := &CacheDefault{Cid: iCid, Data: iData, Expire: iExpire, Created: iCreated, Serialized: iSerialized, Tags: iTags, Checksum: iChecksum}
	if _, err := Engine.Insert(_CacheDefault); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCacheDefault Post CacheDefault via iCacheDefault
func PostCacheDefault(iCacheDefault *CacheDefault) (string, error) {
	_, err := Engine.Insert(iCacheDefault)
	return iCacheDefault.Cid, err
}

// PutCacheDefault Put CacheDefault
func PutCacheDefault(iCacheDefault *CacheDefault) (string, error) {
	_, err := Engine.Id(iCacheDefault.Cid).Update(iCacheDefault)
	return iCacheDefault.Cid, err
}

// PutCacheDefaultViaCid Put CacheDefault via Cid
func PutCacheDefaultViaCid(Cid_ string, iCacheDefault *CacheDefault) (int64, error) {
	row, err := Engine.Update(iCacheDefault, &CacheDefault{Cid: Cid_})
	return row, err
}

// PutCacheDefaultViaData Put CacheDefault via Data
func PutCacheDefaultViaData(Data_ []byte, iCacheDefault *CacheDefault) (int64, error) {
	row, err := Engine.Update(iCacheDefault, &CacheDefault{Data: Data_})
	return row, err
}

// PutCacheDefaultViaExpire Put CacheDefault via Expire
func PutCacheDefaultViaExpire(Expire_ int, iCacheDefault *CacheDefault) (int64, error) {
	row, err := Engine.Update(iCacheDefault, &CacheDefault{Expire: Expire_})
	return row, err
}

// PutCacheDefaultViaCreated Put CacheDefault via Created
func PutCacheDefaultViaCreated(Created_ string, iCacheDefault *CacheDefault) (int64, error) {
	row, err := Engine.Update(iCacheDefault, &CacheDefault{Created: Created_})
	return row, err
}

// PutCacheDefaultViaSerialized Put CacheDefault via Serialized
func PutCacheDefaultViaSerialized(Serialized_ int, iCacheDefault *CacheDefault) (int64, error) {
	row, err := Engine.Update(iCacheDefault, &CacheDefault{Serialized: Serialized_})
	return row, err
}

// PutCacheDefaultViaTags Put CacheDefault via Tags
func PutCacheDefaultViaTags(Tags_ string, iCacheDefault *CacheDefault) (int64, error) {
	row, err := Engine.Update(iCacheDefault, &CacheDefault{Tags: Tags_})
	return row, err
}

// PutCacheDefaultViaChecksum Put CacheDefault via Checksum
func PutCacheDefaultViaChecksum(Checksum_ string, iCacheDefault *CacheDefault) (int64, error) {
	row, err := Engine.Update(iCacheDefault, &CacheDefault{Checksum: Checksum_})
	return row, err
}

// UpdateCacheDefaultViaCid via map[string]interface{}{}
func UpdateCacheDefaultViaCid(iCid string, iCacheDefaultMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDefault)).Where("cid = ?", iCid).Update(iCacheDefaultMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDefaultViaData via map[string]interface{}{}
func UpdateCacheDefaultViaData(iData []byte, iCacheDefaultMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDefault)).Where("data = ?", iData).Update(iCacheDefaultMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDefaultViaExpire via map[string]interface{}{}
func UpdateCacheDefaultViaExpire(iExpire int, iCacheDefaultMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDefault)).Where("expire = ?", iExpire).Update(iCacheDefaultMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDefaultViaCreated via map[string]interface{}{}
func UpdateCacheDefaultViaCreated(iCreated string, iCacheDefaultMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDefault)).Where("created = ?", iCreated).Update(iCacheDefaultMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDefaultViaSerialized via map[string]interface{}{}
func UpdateCacheDefaultViaSerialized(iSerialized int, iCacheDefaultMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDefault)).Where("serialized = ?", iSerialized).Update(iCacheDefaultMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDefaultViaTags via map[string]interface{}{}
func UpdateCacheDefaultViaTags(iTags string, iCacheDefaultMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDefault)).Where("tags = ?", iTags).Update(iCacheDefaultMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheDefaultViaChecksum via map[string]interface{}{}
func UpdateCacheDefaultViaChecksum(iChecksum string, iCacheDefaultMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheDefault)).Where("checksum = ?", iChecksum).Update(iCacheDefaultMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCacheDefault Delete CacheDefault
func DeleteCacheDefault(iCid string) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(CacheDefault))
	return row, err
}

// DeleteCacheDefaultViaCid Delete CacheDefault via Cid
func DeleteCacheDefaultViaCid(iCid string) (err error) {
	var has bool
	var _CacheDefault = &CacheDefault{Cid: iCid}
	if has, err = Engine.Get(_CacheDefault); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CacheDefault)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDefaultViaData Delete CacheDefault via Data
func DeleteCacheDefaultViaData(iData []byte) (err error) {
	var has bool
	var _CacheDefault = &CacheDefault{Data: iData}
	if has, err = Engine.Get(_CacheDefault); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(CacheDefault)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDefaultViaExpire Delete CacheDefault via Expire
func DeleteCacheDefaultViaExpire(iExpire int) (err error) {
	var has bool
	var _CacheDefault = &CacheDefault{Expire: iExpire}
	if has, err = Engine.Get(_CacheDefault); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(CacheDefault)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDefaultViaCreated Delete CacheDefault via Created
func DeleteCacheDefaultViaCreated(iCreated string) (err error) {
	var has bool
	var _CacheDefault = &CacheDefault{Created: iCreated}
	if has, err = Engine.Get(_CacheDefault); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(CacheDefault)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDefaultViaSerialized Delete CacheDefault via Serialized
func DeleteCacheDefaultViaSerialized(iSerialized int) (err error) {
	var has bool
	var _CacheDefault = &CacheDefault{Serialized: iSerialized}
	if has, err = Engine.Get(_CacheDefault); (has == true) && (err == nil) {
		if row, err := Engine.Where("serialized = ?", iSerialized).Delete(new(CacheDefault)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDefaultViaTags Delete CacheDefault via Tags
func DeleteCacheDefaultViaTags(iTags string) (err error) {
	var has bool
	var _CacheDefault = &CacheDefault{Tags: iTags}
	if has, err = Engine.Get(_CacheDefault); (has == true) && (err == nil) {
		if row, err := Engine.Where("tags = ?", iTags).Delete(new(CacheDefault)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheDefaultViaChecksum Delete CacheDefault via Checksum
func DeleteCacheDefaultViaChecksum(iChecksum string) (err error) {
	var has bool
	var _CacheDefault = &CacheDefault{Checksum: iChecksum}
	if has, err = Engine.Get(_CacheDefault); (has == true) && (err == nil) {
		if row, err := Engine.Where("checksum = ?", iChecksum).Delete(new(CacheDefault)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
