package model

type CacheBootstrap struct {
	Cid        string `xorm:"not null pk default '' VARCHAR(255)"`
	Data       []byte `xorm:"LONGBLOB"`
	Expire     int    `xorm:"not null default 0 index INT(11)"`
	Created    string `xorm:"not null default 0.000 DECIMAL(14,3)"`
	Serialized int    `xorm:"not null default 0 SMALLINT(6)"`
	Tags       string `xorm:"LONGTEXT"`
	Checksum   string `xorm:"not null VARCHAR(255)"`
}

// GetCacheBootstrapsCount CacheBootstraps' Count
func GetCacheBootstrapsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CacheBootstrap{})
	return total, err
}

// GetCacheBootstrapCountViaCid Get CacheBootstrap via Cid
func GetCacheBootstrapCountViaCid(iCid string) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CacheBootstrap{Cid: iCid})
	return n
}

// GetCacheBootstrapCountViaData Get CacheBootstrap via Data
func GetCacheBootstrapCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&CacheBootstrap{Data: iData})
	return n
}

// GetCacheBootstrapCountViaExpire Get CacheBootstrap via Expire
func GetCacheBootstrapCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&CacheBootstrap{Expire: iExpire})
	return n
}

// GetCacheBootstrapCountViaCreated Get CacheBootstrap via Created
func GetCacheBootstrapCountViaCreated(iCreated string) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&CacheBootstrap{Created: iCreated})
	return n
}

// GetCacheBootstrapCountViaSerialized Get CacheBootstrap via Serialized
func GetCacheBootstrapCountViaSerialized(iSerialized int) int64 {
	n, _ := Engine.Where("serialized = ?", iSerialized).Count(&CacheBootstrap{Serialized: iSerialized})
	return n
}

// GetCacheBootstrapCountViaTags Get CacheBootstrap via Tags
func GetCacheBootstrapCountViaTags(iTags string) int64 {
	n, _ := Engine.Where("tags = ?", iTags).Count(&CacheBootstrap{Tags: iTags})
	return n
}

// GetCacheBootstrapCountViaChecksum Get CacheBootstrap via Checksum
func GetCacheBootstrapCountViaChecksum(iChecksum string) int64 {
	n, _ := Engine.Where("checksum = ?", iChecksum).Count(&CacheBootstrap{Checksum: iChecksum})
	return n
}

// GetCacheBootstrapsByCidAndDataAndExpire Get CacheBootstraps via CidAndDataAndExpire
func GetCacheBootstrapsByCidAndDataAndExpire(offset int, limit int, Cid_ string, Data_ []byte, Expire_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and data = ? and expire = ?", Cid_, Data_, Expire_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndDataAndCreated Get CacheBootstraps via CidAndDataAndCreated
func GetCacheBootstrapsByCidAndDataAndCreated(offset int, limit int, Cid_ string, Data_ []byte, Created_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and data = ? and created = ?", Cid_, Data_, Created_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndDataAndSerialized Get CacheBootstraps via CidAndDataAndSerialized
func GetCacheBootstrapsByCidAndDataAndSerialized(offset int, limit int, Cid_ string, Data_ []byte, Serialized_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and data = ? and serialized = ?", Cid_, Data_, Serialized_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndDataAndTags Get CacheBootstraps via CidAndDataAndTags
func GetCacheBootstrapsByCidAndDataAndTags(offset int, limit int, Cid_ string, Data_ []byte, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and data = ? and tags = ?", Cid_, Data_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndDataAndChecksum Get CacheBootstraps via CidAndDataAndChecksum
func GetCacheBootstrapsByCidAndDataAndChecksum(offset int, limit int, Cid_ string, Data_ []byte, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and data = ? and checksum = ?", Cid_, Data_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndExpireAndCreated Get CacheBootstraps via CidAndExpireAndCreated
func GetCacheBootstrapsByCidAndExpireAndCreated(offset int, limit int, Cid_ string, Expire_ int, Created_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and expire = ? and created = ?", Cid_, Expire_, Created_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndExpireAndSerialized Get CacheBootstraps via CidAndExpireAndSerialized
func GetCacheBootstrapsByCidAndExpireAndSerialized(offset int, limit int, Cid_ string, Expire_ int, Serialized_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and expire = ? and serialized = ?", Cid_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndExpireAndTags Get CacheBootstraps via CidAndExpireAndTags
func GetCacheBootstrapsByCidAndExpireAndTags(offset int, limit int, Cid_ string, Expire_ int, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and expire = ? and tags = ?", Cid_, Expire_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndExpireAndChecksum Get CacheBootstraps via CidAndExpireAndChecksum
func GetCacheBootstrapsByCidAndExpireAndChecksum(offset int, limit int, Cid_ string, Expire_ int, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and expire = ? and checksum = ?", Cid_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndCreatedAndSerialized Get CacheBootstraps via CidAndCreatedAndSerialized
func GetCacheBootstrapsByCidAndCreatedAndSerialized(offset int, limit int, Cid_ string, Created_ string, Serialized_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and created = ? and serialized = ?", Cid_, Created_, Serialized_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndCreatedAndTags Get CacheBootstraps via CidAndCreatedAndTags
func GetCacheBootstrapsByCidAndCreatedAndTags(offset int, limit int, Cid_ string, Created_ string, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and created = ? and tags = ?", Cid_, Created_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndCreatedAndChecksum Get CacheBootstraps via CidAndCreatedAndChecksum
func GetCacheBootstrapsByCidAndCreatedAndChecksum(offset int, limit int, Cid_ string, Created_ string, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and created = ? and checksum = ?", Cid_, Created_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndSerializedAndTags Get CacheBootstraps via CidAndSerializedAndTags
func GetCacheBootstrapsByCidAndSerializedAndTags(offset int, limit int, Cid_ string, Serialized_ int, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and serialized = ? and tags = ?", Cid_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndSerializedAndChecksum Get CacheBootstraps via CidAndSerializedAndChecksum
func GetCacheBootstrapsByCidAndSerializedAndChecksum(offset int, limit int, Cid_ string, Serialized_ int, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and serialized = ? and checksum = ?", Cid_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndTagsAndChecksum Get CacheBootstraps via CidAndTagsAndChecksum
func GetCacheBootstrapsByCidAndTagsAndChecksum(offset int, limit int, Cid_ string, Tags_ string, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and tags = ? and checksum = ?", Cid_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndExpireAndCreated Get CacheBootstraps via DataAndExpireAndCreated
func GetCacheBootstrapsByDataAndExpireAndCreated(offset int, limit int, Data_ []byte, Expire_ int, Created_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and expire = ? and created = ?", Data_, Expire_, Created_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndExpireAndSerialized Get CacheBootstraps via DataAndExpireAndSerialized
func GetCacheBootstrapsByDataAndExpireAndSerialized(offset int, limit int, Data_ []byte, Expire_ int, Serialized_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and expire = ? and serialized = ?", Data_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndExpireAndTags Get CacheBootstraps via DataAndExpireAndTags
func GetCacheBootstrapsByDataAndExpireAndTags(offset int, limit int, Data_ []byte, Expire_ int, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and expire = ? and tags = ?", Data_, Expire_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndExpireAndChecksum Get CacheBootstraps via DataAndExpireAndChecksum
func GetCacheBootstrapsByDataAndExpireAndChecksum(offset int, limit int, Data_ []byte, Expire_ int, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and expire = ? and checksum = ?", Data_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndCreatedAndSerialized Get CacheBootstraps via DataAndCreatedAndSerialized
func GetCacheBootstrapsByDataAndCreatedAndSerialized(offset int, limit int, Data_ []byte, Created_ string, Serialized_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and created = ? and serialized = ?", Data_, Created_, Serialized_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndCreatedAndTags Get CacheBootstraps via DataAndCreatedAndTags
func GetCacheBootstrapsByDataAndCreatedAndTags(offset int, limit int, Data_ []byte, Created_ string, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and created = ? and tags = ?", Data_, Created_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndCreatedAndChecksum Get CacheBootstraps via DataAndCreatedAndChecksum
func GetCacheBootstrapsByDataAndCreatedAndChecksum(offset int, limit int, Data_ []byte, Created_ string, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and created = ? and checksum = ?", Data_, Created_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndSerializedAndTags Get CacheBootstraps via DataAndSerializedAndTags
func GetCacheBootstrapsByDataAndSerializedAndTags(offset int, limit int, Data_ []byte, Serialized_ int, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and serialized = ? and tags = ?", Data_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndSerializedAndChecksum Get CacheBootstraps via DataAndSerializedAndChecksum
func GetCacheBootstrapsByDataAndSerializedAndChecksum(offset int, limit int, Data_ []byte, Serialized_ int, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and serialized = ? and checksum = ?", Data_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndTagsAndChecksum Get CacheBootstraps via DataAndTagsAndChecksum
func GetCacheBootstrapsByDataAndTagsAndChecksum(offset int, limit int, Data_ []byte, Tags_ string, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and tags = ? and checksum = ?", Data_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByExpireAndCreatedAndSerialized Get CacheBootstraps via ExpireAndCreatedAndSerialized
func GetCacheBootstrapsByExpireAndCreatedAndSerialized(offset int, limit int, Expire_ int, Created_ string, Serialized_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("expire = ? and created = ? and serialized = ?", Expire_, Created_, Serialized_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByExpireAndCreatedAndTags Get CacheBootstraps via ExpireAndCreatedAndTags
func GetCacheBootstrapsByExpireAndCreatedAndTags(offset int, limit int, Expire_ int, Created_ string, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("expire = ? and created = ? and tags = ?", Expire_, Created_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByExpireAndCreatedAndChecksum Get CacheBootstraps via ExpireAndCreatedAndChecksum
func GetCacheBootstrapsByExpireAndCreatedAndChecksum(offset int, limit int, Expire_ int, Created_ string, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("expire = ? and created = ? and checksum = ?", Expire_, Created_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByExpireAndSerializedAndTags Get CacheBootstraps via ExpireAndSerializedAndTags
func GetCacheBootstrapsByExpireAndSerializedAndTags(offset int, limit int, Expire_ int, Serialized_ int, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("expire = ? and serialized = ? and tags = ?", Expire_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByExpireAndSerializedAndChecksum Get CacheBootstraps via ExpireAndSerializedAndChecksum
func GetCacheBootstrapsByExpireAndSerializedAndChecksum(offset int, limit int, Expire_ int, Serialized_ int, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("expire = ? and serialized = ? and checksum = ?", Expire_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByExpireAndTagsAndChecksum Get CacheBootstraps via ExpireAndTagsAndChecksum
func GetCacheBootstrapsByExpireAndTagsAndChecksum(offset int, limit int, Expire_ int, Tags_ string, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("expire = ? and tags = ? and checksum = ?", Expire_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCreatedAndSerializedAndTags Get CacheBootstraps via CreatedAndSerializedAndTags
func GetCacheBootstrapsByCreatedAndSerializedAndTags(offset int, limit int, Created_ string, Serialized_ int, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("created = ? and serialized = ? and tags = ?", Created_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCreatedAndSerializedAndChecksum Get CacheBootstraps via CreatedAndSerializedAndChecksum
func GetCacheBootstrapsByCreatedAndSerializedAndChecksum(offset int, limit int, Created_ string, Serialized_ int, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("created = ? and serialized = ? and checksum = ?", Created_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCreatedAndTagsAndChecksum Get CacheBootstraps via CreatedAndTagsAndChecksum
func GetCacheBootstrapsByCreatedAndTagsAndChecksum(offset int, limit int, Created_ string, Tags_ string, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("created = ? and tags = ? and checksum = ?", Created_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsBySerializedAndTagsAndChecksum Get CacheBootstraps via SerializedAndTagsAndChecksum
func GetCacheBootstrapsBySerializedAndTagsAndChecksum(offset int, limit int, Serialized_ int, Tags_ string, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("serialized = ? and tags = ? and checksum = ?", Serialized_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndData Get CacheBootstraps via CidAndData
func GetCacheBootstrapsByCidAndData(offset int, limit int, Cid_ string, Data_ []byte) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and data = ?", Cid_, Data_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndExpire Get CacheBootstraps via CidAndExpire
func GetCacheBootstrapsByCidAndExpire(offset int, limit int, Cid_ string, Expire_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and expire = ?", Cid_, Expire_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndCreated Get CacheBootstraps via CidAndCreated
func GetCacheBootstrapsByCidAndCreated(offset int, limit int, Cid_ string, Created_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and created = ?", Cid_, Created_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndSerialized Get CacheBootstraps via CidAndSerialized
func GetCacheBootstrapsByCidAndSerialized(offset int, limit int, Cid_ string, Serialized_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and serialized = ?", Cid_, Serialized_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndTags Get CacheBootstraps via CidAndTags
func GetCacheBootstrapsByCidAndTags(offset int, limit int, Cid_ string, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and tags = ?", Cid_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCidAndChecksum Get CacheBootstraps via CidAndChecksum
func GetCacheBootstrapsByCidAndChecksum(offset int, limit int, Cid_ string, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ? and checksum = ?", Cid_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndExpire Get CacheBootstraps via DataAndExpire
func GetCacheBootstrapsByDataAndExpire(offset int, limit int, Data_ []byte, Expire_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and expire = ?", Data_, Expire_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndCreated Get CacheBootstraps via DataAndCreated
func GetCacheBootstrapsByDataAndCreated(offset int, limit int, Data_ []byte, Created_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and created = ?", Data_, Created_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndSerialized Get CacheBootstraps via DataAndSerialized
func GetCacheBootstrapsByDataAndSerialized(offset int, limit int, Data_ []byte, Serialized_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and serialized = ?", Data_, Serialized_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndTags Get CacheBootstraps via DataAndTags
func GetCacheBootstrapsByDataAndTags(offset int, limit int, Data_ []byte, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and tags = ?", Data_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByDataAndChecksum Get CacheBootstraps via DataAndChecksum
func GetCacheBootstrapsByDataAndChecksum(offset int, limit int, Data_ []byte, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ? and checksum = ?", Data_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByExpireAndCreated Get CacheBootstraps via ExpireAndCreated
func GetCacheBootstrapsByExpireAndCreated(offset int, limit int, Expire_ int, Created_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("expire = ? and created = ?", Expire_, Created_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByExpireAndSerialized Get CacheBootstraps via ExpireAndSerialized
func GetCacheBootstrapsByExpireAndSerialized(offset int, limit int, Expire_ int, Serialized_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("expire = ? and serialized = ?", Expire_, Serialized_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByExpireAndTags Get CacheBootstraps via ExpireAndTags
func GetCacheBootstrapsByExpireAndTags(offset int, limit int, Expire_ int, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("expire = ? and tags = ?", Expire_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByExpireAndChecksum Get CacheBootstraps via ExpireAndChecksum
func GetCacheBootstrapsByExpireAndChecksum(offset int, limit int, Expire_ int, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("expire = ? and checksum = ?", Expire_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCreatedAndSerialized Get CacheBootstraps via CreatedAndSerialized
func GetCacheBootstrapsByCreatedAndSerialized(offset int, limit int, Created_ string, Serialized_ int) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("created = ? and serialized = ?", Created_, Serialized_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCreatedAndTags Get CacheBootstraps via CreatedAndTags
func GetCacheBootstrapsByCreatedAndTags(offset int, limit int, Created_ string, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("created = ? and tags = ?", Created_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByCreatedAndChecksum Get CacheBootstraps via CreatedAndChecksum
func GetCacheBootstrapsByCreatedAndChecksum(offset int, limit int, Created_ string, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("created = ? and checksum = ?", Created_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsBySerializedAndTags Get CacheBootstraps via SerializedAndTags
func GetCacheBootstrapsBySerializedAndTags(offset int, limit int, Serialized_ int, Tags_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("serialized = ? and tags = ?", Serialized_, Tags_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsBySerializedAndChecksum Get CacheBootstraps via SerializedAndChecksum
func GetCacheBootstrapsBySerializedAndChecksum(offset int, limit int, Serialized_ int, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("serialized = ? and checksum = ?", Serialized_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsByTagsAndChecksum Get CacheBootstraps via TagsAndChecksum
func GetCacheBootstrapsByTagsAndChecksum(offset int, limit int, Tags_ string, Checksum_ string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("tags = ? and checksum = ?", Tags_, Checksum_).Limit(limit, offset).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstraps Get CacheBootstraps via field
func GetCacheBootstraps(offset int, limit int, field string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Limit(limit, offset).Desc(field).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsViaCid Get CacheBootstraps via Cid
func GetCacheBootstrapsViaCid(offset int, limit int, Cid_ string, field string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsViaData Get CacheBootstraps via Data
func GetCacheBootstrapsViaData(offset int, limit int, Data_ []byte, field string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsViaExpire Get CacheBootstraps via Expire
func GetCacheBootstrapsViaExpire(offset int, limit int, Expire_ int, field string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsViaCreated Get CacheBootstraps via Created
func GetCacheBootstrapsViaCreated(offset int, limit int, Created_ string, field string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsViaSerialized Get CacheBootstraps via Serialized
func GetCacheBootstrapsViaSerialized(offset int, limit int, Serialized_ int, field string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("serialized = ?", Serialized_).Limit(limit, offset).Desc(field).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsViaTags Get CacheBootstraps via Tags
func GetCacheBootstrapsViaTags(offset int, limit int, Tags_ string, field string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("tags = ?", Tags_).Limit(limit, offset).Desc(field).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// GetCacheBootstrapsViaChecksum Get CacheBootstraps via Checksum
func GetCacheBootstrapsViaChecksum(offset int, limit int, Checksum_ string, field string) (*[]*CacheBootstrap, error) {
	var _CacheBootstrap = new([]*CacheBootstrap)
	err := Engine.Table("cache_bootstrap").Where("checksum = ?", Checksum_).Limit(limit, offset).Desc(field).Find(_CacheBootstrap)
	return _CacheBootstrap, err
}

// HasCacheBootstrapViaCid Has CacheBootstrap via Cid
func HasCacheBootstrapViaCid(iCid string) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CacheBootstrap)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheBootstrapViaData Has CacheBootstrap via Data
func HasCacheBootstrapViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(CacheBootstrap)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheBootstrapViaExpire Has CacheBootstrap via Expire
func HasCacheBootstrapViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(CacheBootstrap)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheBootstrapViaCreated Has CacheBootstrap via Created
func HasCacheBootstrapViaCreated(iCreated string) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(CacheBootstrap)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheBootstrapViaSerialized Has CacheBootstrap via Serialized
func HasCacheBootstrapViaSerialized(iSerialized int) bool {
	if has, err := Engine.Where("serialized = ?", iSerialized).Get(new(CacheBootstrap)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheBootstrapViaTags Has CacheBootstrap via Tags
func HasCacheBootstrapViaTags(iTags string) bool {
	if has, err := Engine.Where("tags = ?", iTags).Get(new(CacheBootstrap)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheBootstrapViaChecksum Has CacheBootstrap via Checksum
func HasCacheBootstrapViaChecksum(iChecksum string) bool {
	if has, err := Engine.Where("checksum = ?", iChecksum).Get(new(CacheBootstrap)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCacheBootstrapViaCid Get CacheBootstrap via Cid
func GetCacheBootstrapViaCid(iCid string) (*CacheBootstrap, error) {
	var _CacheBootstrap = &CacheBootstrap{Cid: iCid}
	has, err := Engine.Get(_CacheBootstrap)
	if has {
		return _CacheBootstrap, err
	} else {
		return nil, err
	}
}

// GetCacheBootstrapViaData Get CacheBootstrap via Data
func GetCacheBootstrapViaData(iData []byte) (*CacheBootstrap, error) {
	var _CacheBootstrap = &CacheBootstrap{Data: iData}
	has, err := Engine.Get(_CacheBootstrap)
	if has {
		return _CacheBootstrap, err
	} else {
		return nil, err
	}
}

// GetCacheBootstrapViaExpire Get CacheBootstrap via Expire
func GetCacheBootstrapViaExpire(iExpire int) (*CacheBootstrap, error) {
	var _CacheBootstrap = &CacheBootstrap{Expire: iExpire}
	has, err := Engine.Get(_CacheBootstrap)
	if has {
		return _CacheBootstrap, err
	} else {
		return nil, err
	}
}

// GetCacheBootstrapViaCreated Get CacheBootstrap via Created
func GetCacheBootstrapViaCreated(iCreated string) (*CacheBootstrap, error) {
	var _CacheBootstrap = &CacheBootstrap{Created: iCreated}
	has, err := Engine.Get(_CacheBootstrap)
	if has {
		return _CacheBootstrap, err
	} else {
		return nil, err
	}
}

// GetCacheBootstrapViaSerialized Get CacheBootstrap via Serialized
func GetCacheBootstrapViaSerialized(iSerialized int) (*CacheBootstrap, error) {
	var _CacheBootstrap = &CacheBootstrap{Serialized: iSerialized}
	has, err := Engine.Get(_CacheBootstrap)
	if has {
		return _CacheBootstrap, err
	} else {
		return nil, err
	}
}

// GetCacheBootstrapViaTags Get CacheBootstrap via Tags
func GetCacheBootstrapViaTags(iTags string) (*CacheBootstrap, error) {
	var _CacheBootstrap = &CacheBootstrap{Tags: iTags}
	has, err := Engine.Get(_CacheBootstrap)
	if has {
		return _CacheBootstrap, err
	} else {
		return nil, err
	}
}

// GetCacheBootstrapViaChecksum Get CacheBootstrap via Checksum
func GetCacheBootstrapViaChecksum(iChecksum string) (*CacheBootstrap, error) {
	var _CacheBootstrap = &CacheBootstrap{Checksum: iChecksum}
	has, err := Engine.Get(_CacheBootstrap)
	if has {
		return _CacheBootstrap, err
	} else {
		return nil, err
	}
}

// SetCacheBootstrapViaCid Set CacheBootstrap via Cid
func SetCacheBootstrapViaCid(iCid string, cache_bootstrap *CacheBootstrap) (int64, error) {
	cache_bootstrap.Cid = iCid
	return Engine.Insert(cache_bootstrap)
}

// SetCacheBootstrapViaData Set CacheBootstrap via Data
func SetCacheBootstrapViaData(iData []byte, cache_bootstrap *CacheBootstrap) (int64, error) {
	cache_bootstrap.Data = iData
	return Engine.Insert(cache_bootstrap)
}

// SetCacheBootstrapViaExpire Set CacheBootstrap via Expire
func SetCacheBootstrapViaExpire(iExpire int, cache_bootstrap *CacheBootstrap) (int64, error) {
	cache_bootstrap.Expire = iExpire
	return Engine.Insert(cache_bootstrap)
}

// SetCacheBootstrapViaCreated Set CacheBootstrap via Created
func SetCacheBootstrapViaCreated(iCreated string, cache_bootstrap *CacheBootstrap) (int64, error) {
	cache_bootstrap.Created = iCreated
	return Engine.Insert(cache_bootstrap)
}

// SetCacheBootstrapViaSerialized Set CacheBootstrap via Serialized
func SetCacheBootstrapViaSerialized(iSerialized int, cache_bootstrap *CacheBootstrap) (int64, error) {
	cache_bootstrap.Serialized = iSerialized
	return Engine.Insert(cache_bootstrap)
}

// SetCacheBootstrapViaTags Set CacheBootstrap via Tags
func SetCacheBootstrapViaTags(iTags string, cache_bootstrap *CacheBootstrap) (int64, error) {
	cache_bootstrap.Tags = iTags
	return Engine.Insert(cache_bootstrap)
}

// SetCacheBootstrapViaChecksum Set CacheBootstrap via Checksum
func SetCacheBootstrapViaChecksum(iChecksum string, cache_bootstrap *CacheBootstrap) (int64, error) {
	cache_bootstrap.Checksum = iChecksum
	return Engine.Insert(cache_bootstrap)
}

// AddCacheBootstrap Add CacheBootstrap via all columns
func AddCacheBootstrap(iCid string, iData []byte, iExpire int, iCreated string, iSerialized int, iTags string, iChecksum string) error {
	_CacheBootstrap := &CacheBootstrap{Cid: iCid, Data: iData, Expire: iExpire, Created: iCreated, Serialized: iSerialized, Tags: iTags, Checksum: iChecksum}
	if _, err := Engine.Insert(_CacheBootstrap); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCacheBootstrap Post CacheBootstrap via iCacheBootstrap
func PostCacheBootstrap(iCacheBootstrap *CacheBootstrap) (string, error) {
	_, err := Engine.Insert(iCacheBootstrap)
	return iCacheBootstrap.Cid, err
}

// PutCacheBootstrap Put CacheBootstrap
func PutCacheBootstrap(iCacheBootstrap *CacheBootstrap) (string, error) {
	_, err := Engine.Id(iCacheBootstrap.Cid).Update(iCacheBootstrap)
	return iCacheBootstrap.Cid, err
}

// PutCacheBootstrapViaCid Put CacheBootstrap via Cid
func PutCacheBootstrapViaCid(Cid_ string, iCacheBootstrap *CacheBootstrap) (int64, error) {
	row, err := Engine.Update(iCacheBootstrap, &CacheBootstrap{Cid: Cid_})
	return row, err
}

// PutCacheBootstrapViaData Put CacheBootstrap via Data
func PutCacheBootstrapViaData(Data_ []byte, iCacheBootstrap *CacheBootstrap) (int64, error) {
	row, err := Engine.Update(iCacheBootstrap, &CacheBootstrap{Data: Data_})
	return row, err
}

// PutCacheBootstrapViaExpire Put CacheBootstrap via Expire
func PutCacheBootstrapViaExpire(Expire_ int, iCacheBootstrap *CacheBootstrap) (int64, error) {
	row, err := Engine.Update(iCacheBootstrap, &CacheBootstrap{Expire: Expire_})
	return row, err
}

// PutCacheBootstrapViaCreated Put CacheBootstrap via Created
func PutCacheBootstrapViaCreated(Created_ string, iCacheBootstrap *CacheBootstrap) (int64, error) {
	row, err := Engine.Update(iCacheBootstrap, &CacheBootstrap{Created: Created_})
	return row, err
}

// PutCacheBootstrapViaSerialized Put CacheBootstrap via Serialized
func PutCacheBootstrapViaSerialized(Serialized_ int, iCacheBootstrap *CacheBootstrap) (int64, error) {
	row, err := Engine.Update(iCacheBootstrap, &CacheBootstrap{Serialized: Serialized_})
	return row, err
}

// PutCacheBootstrapViaTags Put CacheBootstrap via Tags
func PutCacheBootstrapViaTags(Tags_ string, iCacheBootstrap *CacheBootstrap) (int64, error) {
	row, err := Engine.Update(iCacheBootstrap, &CacheBootstrap{Tags: Tags_})
	return row, err
}

// PutCacheBootstrapViaChecksum Put CacheBootstrap via Checksum
func PutCacheBootstrapViaChecksum(Checksum_ string, iCacheBootstrap *CacheBootstrap) (int64, error) {
	row, err := Engine.Update(iCacheBootstrap, &CacheBootstrap{Checksum: Checksum_})
	return row, err
}

// UpdateCacheBootstrapViaCid via map[string]interface{}{}
func UpdateCacheBootstrapViaCid(iCid string, iCacheBootstrapMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheBootstrap)).Where("cid = ?", iCid).Update(iCacheBootstrapMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheBootstrapViaData via map[string]interface{}{}
func UpdateCacheBootstrapViaData(iData []byte, iCacheBootstrapMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheBootstrap)).Where("data = ?", iData).Update(iCacheBootstrapMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheBootstrapViaExpire via map[string]interface{}{}
func UpdateCacheBootstrapViaExpire(iExpire int, iCacheBootstrapMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheBootstrap)).Where("expire = ?", iExpire).Update(iCacheBootstrapMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheBootstrapViaCreated via map[string]interface{}{}
func UpdateCacheBootstrapViaCreated(iCreated string, iCacheBootstrapMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheBootstrap)).Where("created = ?", iCreated).Update(iCacheBootstrapMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheBootstrapViaSerialized via map[string]interface{}{}
func UpdateCacheBootstrapViaSerialized(iSerialized int, iCacheBootstrapMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheBootstrap)).Where("serialized = ?", iSerialized).Update(iCacheBootstrapMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheBootstrapViaTags via map[string]interface{}{}
func UpdateCacheBootstrapViaTags(iTags string, iCacheBootstrapMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheBootstrap)).Where("tags = ?", iTags).Update(iCacheBootstrapMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheBootstrapViaChecksum via map[string]interface{}{}
func UpdateCacheBootstrapViaChecksum(iChecksum string, iCacheBootstrapMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheBootstrap)).Where("checksum = ?", iChecksum).Update(iCacheBootstrapMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCacheBootstrap Delete CacheBootstrap
func DeleteCacheBootstrap(iCid string) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(CacheBootstrap))
	return row, err
}

// DeleteCacheBootstrapViaCid Delete CacheBootstrap via Cid
func DeleteCacheBootstrapViaCid(iCid string) (err error) {
	var has bool
	var _CacheBootstrap = &CacheBootstrap{Cid: iCid}
	if has, err = Engine.Get(_CacheBootstrap); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CacheBootstrap)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheBootstrapViaData Delete CacheBootstrap via Data
func DeleteCacheBootstrapViaData(iData []byte) (err error) {
	var has bool
	var _CacheBootstrap = &CacheBootstrap{Data: iData}
	if has, err = Engine.Get(_CacheBootstrap); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(CacheBootstrap)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheBootstrapViaExpire Delete CacheBootstrap via Expire
func DeleteCacheBootstrapViaExpire(iExpire int) (err error) {
	var has bool
	var _CacheBootstrap = &CacheBootstrap{Expire: iExpire}
	if has, err = Engine.Get(_CacheBootstrap); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(CacheBootstrap)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheBootstrapViaCreated Delete CacheBootstrap via Created
func DeleteCacheBootstrapViaCreated(iCreated string) (err error) {
	var has bool
	var _CacheBootstrap = &CacheBootstrap{Created: iCreated}
	if has, err = Engine.Get(_CacheBootstrap); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(CacheBootstrap)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheBootstrapViaSerialized Delete CacheBootstrap via Serialized
func DeleteCacheBootstrapViaSerialized(iSerialized int) (err error) {
	var has bool
	var _CacheBootstrap = &CacheBootstrap{Serialized: iSerialized}
	if has, err = Engine.Get(_CacheBootstrap); (has == true) && (err == nil) {
		if row, err := Engine.Where("serialized = ?", iSerialized).Delete(new(CacheBootstrap)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheBootstrapViaTags Delete CacheBootstrap via Tags
func DeleteCacheBootstrapViaTags(iTags string) (err error) {
	var has bool
	var _CacheBootstrap = &CacheBootstrap{Tags: iTags}
	if has, err = Engine.Get(_CacheBootstrap); (has == true) && (err == nil) {
		if row, err := Engine.Where("tags = ?", iTags).Delete(new(CacheBootstrap)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheBootstrapViaChecksum Delete CacheBootstrap via Checksum
func DeleteCacheBootstrapViaChecksum(iChecksum string) (err error) {
	var has bool
	var _CacheBootstrap = &CacheBootstrap{Checksum: iChecksum}
	if has, err = Engine.Get(_CacheBootstrap); (has == true) && (err == nil) {
		if row, err := Engine.Where("checksum = ?", iChecksum).Delete(new(CacheBootstrap)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
