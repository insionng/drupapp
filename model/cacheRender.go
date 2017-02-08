package model

type CacheRender struct {
	Cid        string `xorm:"not null pk default '' VARCHAR(255)"`
	Data       []byte `xorm:"LONGBLOB"`
	Expire     int    `xorm:"not null default 0 index INT(11)"`
	Created    string `xorm:"not null default 0.000 DECIMAL(14,3)"`
	Serialized int    `xorm:"not null default 0 SMALLINT(6)"`
	Tags       string `xorm:"LONGTEXT"`
	Checksum   string `xorm:"not null VARCHAR(255)"`
}

// GetCacheRendersCount CacheRenders' Count
func GetCacheRendersCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CacheRender{})
	return total, err
}

// GetCacheRenderCountViaCid Get CacheRender via Cid
func GetCacheRenderCountViaCid(iCid string) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CacheRender{Cid: iCid})
	return n
}

// GetCacheRenderCountViaData Get CacheRender via Data
func GetCacheRenderCountViaData(iData []byte) int64 {
	n, _ := Engine.Where("data = ?", iData).Count(&CacheRender{Data: iData})
	return n
}

// GetCacheRenderCountViaExpire Get CacheRender via Expire
func GetCacheRenderCountViaExpire(iExpire int) int64 {
	n, _ := Engine.Where("expire = ?", iExpire).Count(&CacheRender{Expire: iExpire})
	return n
}

// GetCacheRenderCountViaCreated Get CacheRender via Created
func GetCacheRenderCountViaCreated(iCreated string) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&CacheRender{Created: iCreated})
	return n
}

// GetCacheRenderCountViaSerialized Get CacheRender via Serialized
func GetCacheRenderCountViaSerialized(iSerialized int) int64 {
	n, _ := Engine.Where("serialized = ?", iSerialized).Count(&CacheRender{Serialized: iSerialized})
	return n
}

// GetCacheRenderCountViaTags Get CacheRender via Tags
func GetCacheRenderCountViaTags(iTags string) int64 {
	n, _ := Engine.Where("tags = ?", iTags).Count(&CacheRender{Tags: iTags})
	return n
}

// GetCacheRenderCountViaChecksum Get CacheRender via Checksum
func GetCacheRenderCountViaChecksum(iChecksum string) int64 {
	n, _ := Engine.Where("checksum = ?", iChecksum).Count(&CacheRender{Checksum: iChecksum})
	return n
}

// GetCacheRendersByCidAndDataAndExpire Get CacheRenders via CidAndDataAndExpire
func GetCacheRendersByCidAndDataAndExpire(offset int, limit int, Cid_ string, Data_ []byte, Expire_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and data = ? and expire = ?", Cid_, Data_, Expire_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndDataAndCreated Get CacheRenders via CidAndDataAndCreated
func GetCacheRendersByCidAndDataAndCreated(offset int, limit int, Cid_ string, Data_ []byte, Created_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and data = ? and created = ?", Cid_, Data_, Created_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndDataAndSerialized Get CacheRenders via CidAndDataAndSerialized
func GetCacheRendersByCidAndDataAndSerialized(offset int, limit int, Cid_ string, Data_ []byte, Serialized_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and data = ? and serialized = ?", Cid_, Data_, Serialized_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndDataAndTags Get CacheRenders via CidAndDataAndTags
func GetCacheRendersByCidAndDataAndTags(offset int, limit int, Cid_ string, Data_ []byte, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and data = ? and tags = ?", Cid_, Data_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndDataAndChecksum Get CacheRenders via CidAndDataAndChecksum
func GetCacheRendersByCidAndDataAndChecksum(offset int, limit int, Cid_ string, Data_ []byte, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and data = ? and checksum = ?", Cid_, Data_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndExpireAndCreated Get CacheRenders via CidAndExpireAndCreated
func GetCacheRendersByCidAndExpireAndCreated(offset int, limit int, Cid_ string, Expire_ int, Created_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and expire = ? and created = ?", Cid_, Expire_, Created_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndExpireAndSerialized Get CacheRenders via CidAndExpireAndSerialized
func GetCacheRendersByCidAndExpireAndSerialized(offset int, limit int, Cid_ string, Expire_ int, Serialized_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and expire = ? and serialized = ?", Cid_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndExpireAndTags Get CacheRenders via CidAndExpireAndTags
func GetCacheRendersByCidAndExpireAndTags(offset int, limit int, Cid_ string, Expire_ int, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and expire = ? and tags = ?", Cid_, Expire_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndExpireAndChecksum Get CacheRenders via CidAndExpireAndChecksum
func GetCacheRendersByCidAndExpireAndChecksum(offset int, limit int, Cid_ string, Expire_ int, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and expire = ? and checksum = ?", Cid_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndCreatedAndSerialized Get CacheRenders via CidAndCreatedAndSerialized
func GetCacheRendersByCidAndCreatedAndSerialized(offset int, limit int, Cid_ string, Created_ string, Serialized_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and created = ? and serialized = ?", Cid_, Created_, Serialized_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndCreatedAndTags Get CacheRenders via CidAndCreatedAndTags
func GetCacheRendersByCidAndCreatedAndTags(offset int, limit int, Cid_ string, Created_ string, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and created = ? and tags = ?", Cid_, Created_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndCreatedAndChecksum Get CacheRenders via CidAndCreatedAndChecksum
func GetCacheRendersByCidAndCreatedAndChecksum(offset int, limit int, Cid_ string, Created_ string, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and created = ? and checksum = ?", Cid_, Created_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndSerializedAndTags Get CacheRenders via CidAndSerializedAndTags
func GetCacheRendersByCidAndSerializedAndTags(offset int, limit int, Cid_ string, Serialized_ int, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and serialized = ? and tags = ?", Cid_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndSerializedAndChecksum Get CacheRenders via CidAndSerializedAndChecksum
func GetCacheRendersByCidAndSerializedAndChecksum(offset int, limit int, Cid_ string, Serialized_ int, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and serialized = ? and checksum = ?", Cid_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndTagsAndChecksum Get CacheRenders via CidAndTagsAndChecksum
func GetCacheRendersByCidAndTagsAndChecksum(offset int, limit int, Cid_ string, Tags_ string, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and tags = ? and checksum = ?", Cid_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndExpireAndCreated Get CacheRenders via DataAndExpireAndCreated
func GetCacheRendersByDataAndExpireAndCreated(offset int, limit int, Data_ []byte, Expire_ int, Created_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and expire = ? and created = ?", Data_, Expire_, Created_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndExpireAndSerialized Get CacheRenders via DataAndExpireAndSerialized
func GetCacheRendersByDataAndExpireAndSerialized(offset int, limit int, Data_ []byte, Expire_ int, Serialized_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and expire = ? and serialized = ?", Data_, Expire_, Serialized_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndExpireAndTags Get CacheRenders via DataAndExpireAndTags
func GetCacheRendersByDataAndExpireAndTags(offset int, limit int, Data_ []byte, Expire_ int, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and expire = ? and tags = ?", Data_, Expire_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndExpireAndChecksum Get CacheRenders via DataAndExpireAndChecksum
func GetCacheRendersByDataAndExpireAndChecksum(offset int, limit int, Data_ []byte, Expire_ int, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and expire = ? and checksum = ?", Data_, Expire_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndCreatedAndSerialized Get CacheRenders via DataAndCreatedAndSerialized
func GetCacheRendersByDataAndCreatedAndSerialized(offset int, limit int, Data_ []byte, Created_ string, Serialized_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and created = ? and serialized = ?", Data_, Created_, Serialized_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndCreatedAndTags Get CacheRenders via DataAndCreatedAndTags
func GetCacheRendersByDataAndCreatedAndTags(offset int, limit int, Data_ []byte, Created_ string, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and created = ? and tags = ?", Data_, Created_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndCreatedAndChecksum Get CacheRenders via DataAndCreatedAndChecksum
func GetCacheRendersByDataAndCreatedAndChecksum(offset int, limit int, Data_ []byte, Created_ string, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and created = ? and checksum = ?", Data_, Created_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndSerializedAndTags Get CacheRenders via DataAndSerializedAndTags
func GetCacheRendersByDataAndSerializedAndTags(offset int, limit int, Data_ []byte, Serialized_ int, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and serialized = ? and tags = ?", Data_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndSerializedAndChecksum Get CacheRenders via DataAndSerializedAndChecksum
func GetCacheRendersByDataAndSerializedAndChecksum(offset int, limit int, Data_ []byte, Serialized_ int, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and serialized = ? and checksum = ?", Data_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndTagsAndChecksum Get CacheRenders via DataAndTagsAndChecksum
func GetCacheRendersByDataAndTagsAndChecksum(offset int, limit int, Data_ []byte, Tags_ string, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and tags = ? and checksum = ?", Data_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByExpireAndCreatedAndSerialized Get CacheRenders via ExpireAndCreatedAndSerialized
func GetCacheRendersByExpireAndCreatedAndSerialized(offset int, limit int, Expire_ int, Created_ string, Serialized_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("expire = ? and created = ? and serialized = ?", Expire_, Created_, Serialized_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByExpireAndCreatedAndTags Get CacheRenders via ExpireAndCreatedAndTags
func GetCacheRendersByExpireAndCreatedAndTags(offset int, limit int, Expire_ int, Created_ string, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("expire = ? and created = ? and tags = ?", Expire_, Created_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByExpireAndCreatedAndChecksum Get CacheRenders via ExpireAndCreatedAndChecksum
func GetCacheRendersByExpireAndCreatedAndChecksum(offset int, limit int, Expire_ int, Created_ string, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("expire = ? and created = ? and checksum = ?", Expire_, Created_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByExpireAndSerializedAndTags Get CacheRenders via ExpireAndSerializedAndTags
func GetCacheRendersByExpireAndSerializedAndTags(offset int, limit int, Expire_ int, Serialized_ int, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("expire = ? and serialized = ? and tags = ?", Expire_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByExpireAndSerializedAndChecksum Get CacheRenders via ExpireAndSerializedAndChecksum
func GetCacheRendersByExpireAndSerializedAndChecksum(offset int, limit int, Expire_ int, Serialized_ int, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("expire = ? and serialized = ? and checksum = ?", Expire_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByExpireAndTagsAndChecksum Get CacheRenders via ExpireAndTagsAndChecksum
func GetCacheRendersByExpireAndTagsAndChecksum(offset int, limit int, Expire_ int, Tags_ string, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("expire = ? and tags = ? and checksum = ?", Expire_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCreatedAndSerializedAndTags Get CacheRenders via CreatedAndSerializedAndTags
func GetCacheRendersByCreatedAndSerializedAndTags(offset int, limit int, Created_ string, Serialized_ int, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("created = ? and serialized = ? and tags = ?", Created_, Serialized_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCreatedAndSerializedAndChecksum Get CacheRenders via CreatedAndSerializedAndChecksum
func GetCacheRendersByCreatedAndSerializedAndChecksum(offset int, limit int, Created_ string, Serialized_ int, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("created = ? and serialized = ? and checksum = ?", Created_, Serialized_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCreatedAndTagsAndChecksum Get CacheRenders via CreatedAndTagsAndChecksum
func GetCacheRendersByCreatedAndTagsAndChecksum(offset int, limit int, Created_ string, Tags_ string, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("created = ? and tags = ? and checksum = ?", Created_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersBySerializedAndTagsAndChecksum Get CacheRenders via SerializedAndTagsAndChecksum
func GetCacheRendersBySerializedAndTagsAndChecksum(offset int, limit int, Serialized_ int, Tags_ string, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("serialized = ? and tags = ? and checksum = ?", Serialized_, Tags_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndData Get CacheRenders via CidAndData
func GetCacheRendersByCidAndData(offset int, limit int, Cid_ string, Data_ []byte) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and data = ?", Cid_, Data_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndExpire Get CacheRenders via CidAndExpire
func GetCacheRendersByCidAndExpire(offset int, limit int, Cid_ string, Expire_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and expire = ?", Cid_, Expire_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndCreated Get CacheRenders via CidAndCreated
func GetCacheRendersByCidAndCreated(offset int, limit int, Cid_ string, Created_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and created = ?", Cid_, Created_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndSerialized Get CacheRenders via CidAndSerialized
func GetCacheRendersByCidAndSerialized(offset int, limit int, Cid_ string, Serialized_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and serialized = ?", Cid_, Serialized_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndTags Get CacheRenders via CidAndTags
func GetCacheRendersByCidAndTags(offset int, limit int, Cid_ string, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and tags = ?", Cid_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCidAndChecksum Get CacheRenders via CidAndChecksum
func GetCacheRendersByCidAndChecksum(offset int, limit int, Cid_ string, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ? and checksum = ?", Cid_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndExpire Get CacheRenders via DataAndExpire
func GetCacheRendersByDataAndExpire(offset int, limit int, Data_ []byte, Expire_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and expire = ?", Data_, Expire_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndCreated Get CacheRenders via DataAndCreated
func GetCacheRendersByDataAndCreated(offset int, limit int, Data_ []byte, Created_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and created = ?", Data_, Created_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndSerialized Get CacheRenders via DataAndSerialized
func GetCacheRendersByDataAndSerialized(offset int, limit int, Data_ []byte, Serialized_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and serialized = ?", Data_, Serialized_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndTags Get CacheRenders via DataAndTags
func GetCacheRendersByDataAndTags(offset int, limit int, Data_ []byte, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and tags = ?", Data_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByDataAndChecksum Get CacheRenders via DataAndChecksum
func GetCacheRendersByDataAndChecksum(offset int, limit int, Data_ []byte, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ? and checksum = ?", Data_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByExpireAndCreated Get CacheRenders via ExpireAndCreated
func GetCacheRendersByExpireAndCreated(offset int, limit int, Expire_ int, Created_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("expire = ? and created = ?", Expire_, Created_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByExpireAndSerialized Get CacheRenders via ExpireAndSerialized
func GetCacheRendersByExpireAndSerialized(offset int, limit int, Expire_ int, Serialized_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("expire = ? and serialized = ?", Expire_, Serialized_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByExpireAndTags Get CacheRenders via ExpireAndTags
func GetCacheRendersByExpireAndTags(offset int, limit int, Expire_ int, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("expire = ? and tags = ?", Expire_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByExpireAndChecksum Get CacheRenders via ExpireAndChecksum
func GetCacheRendersByExpireAndChecksum(offset int, limit int, Expire_ int, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("expire = ? and checksum = ?", Expire_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCreatedAndSerialized Get CacheRenders via CreatedAndSerialized
func GetCacheRendersByCreatedAndSerialized(offset int, limit int, Created_ string, Serialized_ int) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("created = ? and serialized = ?", Created_, Serialized_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCreatedAndTags Get CacheRenders via CreatedAndTags
func GetCacheRendersByCreatedAndTags(offset int, limit int, Created_ string, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("created = ? and tags = ?", Created_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByCreatedAndChecksum Get CacheRenders via CreatedAndChecksum
func GetCacheRendersByCreatedAndChecksum(offset int, limit int, Created_ string, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("created = ? and checksum = ?", Created_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersBySerializedAndTags Get CacheRenders via SerializedAndTags
func GetCacheRendersBySerializedAndTags(offset int, limit int, Serialized_ int, Tags_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("serialized = ? and tags = ?", Serialized_, Tags_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersBySerializedAndChecksum Get CacheRenders via SerializedAndChecksum
func GetCacheRendersBySerializedAndChecksum(offset int, limit int, Serialized_ int, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("serialized = ? and checksum = ?", Serialized_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersByTagsAndChecksum Get CacheRenders via TagsAndChecksum
func GetCacheRendersByTagsAndChecksum(offset int, limit int, Tags_ string, Checksum_ string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("tags = ? and checksum = ?", Tags_, Checksum_).Limit(limit, offset).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRenders Get CacheRenders via field
func GetCacheRenders(offset int, limit int, field string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Limit(limit, offset).Desc(field).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersViaCid Get CacheRenders via Cid
func GetCacheRendersViaCid(offset int, limit int, Cid_ string, field string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersViaData Get CacheRenders via Data
func GetCacheRendersViaData(offset int, limit int, Data_ []byte, field string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("data = ?", Data_).Limit(limit, offset).Desc(field).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersViaExpire Get CacheRenders via Expire
func GetCacheRendersViaExpire(offset int, limit int, Expire_ int, field string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("expire = ?", Expire_).Limit(limit, offset).Desc(field).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersViaCreated Get CacheRenders via Created
func GetCacheRendersViaCreated(offset int, limit int, Created_ string, field string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersViaSerialized Get CacheRenders via Serialized
func GetCacheRendersViaSerialized(offset int, limit int, Serialized_ int, field string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("serialized = ?", Serialized_).Limit(limit, offset).Desc(field).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersViaTags Get CacheRenders via Tags
func GetCacheRendersViaTags(offset int, limit int, Tags_ string, field string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("tags = ?", Tags_).Limit(limit, offset).Desc(field).Find(_CacheRender)
	return _CacheRender, err
}

// GetCacheRendersViaChecksum Get CacheRenders via Checksum
func GetCacheRendersViaChecksum(offset int, limit int, Checksum_ string, field string) (*[]*CacheRender, error) {
	var _CacheRender = new([]*CacheRender)
	err := Engine.Table("cache_render").Where("checksum = ?", Checksum_).Limit(limit, offset).Desc(field).Find(_CacheRender)
	return _CacheRender, err
}

// HasCacheRenderViaCid Has CacheRender via Cid
func HasCacheRenderViaCid(iCid string) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CacheRender)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheRenderViaData Has CacheRender via Data
func HasCacheRenderViaData(iData []byte) bool {
	if has, err := Engine.Where("data = ?", iData).Get(new(CacheRender)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheRenderViaExpire Has CacheRender via Expire
func HasCacheRenderViaExpire(iExpire int) bool {
	if has, err := Engine.Where("expire = ?", iExpire).Get(new(CacheRender)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheRenderViaCreated Has CacheRender via Created
func HasCacheRenderViaCreated(iCreated string) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(CacheRender)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheRenderViaSerialized Has CacheRender via Serialized
func HasCacheRenderViaSerialized(iSerialized int) bool {
	if has, err := Engine.Where("serialized = ?", iSerialized).Get(new(CacheRender)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheRenderViaTags Has CacheRender via Tags
func HasCacheRenderViaTags(iTags string) bool {
	if has, err := Engine.Where("tags = ?", iTags).Get(new(CacheRender)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCacheRenderViaChecksum Has CacheRender via Checksum
func HasCacheRenderViaChecksum(iChecksum string) bool {
	if has, err := Engine.Where("checksum = ?", iChecksum).Get(new(CacheRender)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCacheRenderViaCid Get CacheRender via Cid
func GetCacheRenderViaCid(iCid string) (*CacheRender, error) {
	var _CacheRender = &CacheRender{Cid: iCid}
	has, err := Engine.Get(_CacheRender)
	if has {
		return _CacheRender, err
	} else {
		return nil, err
	}
}

// GetCacheRenderViaData Get CacheRender via Data
func GetCacheRenderViaData(iData []byte) (*CacheRender, error) {
	var _CacheRender = &CacheRender{Data: iData}
	has, err := Engine.Get(_CacheRender)
	if has {
		return _CacheRender, err
	} else {
		return nil, err
	}
}

// GetCacheRenderViaExpire Get CacheRender via Expire
func GetCacheRenderViaExpire(iExpire int) (*CacheRender, error) {
	var _CacheRender = &CacheRender{Expire: iExpire}
	has, err := Engine.Get(_CacheRender)
	if has {
		return _CacheRender, err
	} else {
		return nil, err
	}
}

// GetCacheRenderViaCreated Get CacheRender via Created
func GetCacheRenderViaCreated(iCreated string) (*CacheRender, error) {
	var _CacheRender = &CacheRender{Created: iCreated}
	has, err := Engine.Get(_CacheRender)
	if has {
		return _CacheRender, err
	} else {
		return nil, err
	}
}

// GetCacheRenderViaSerialized Get CacheRender via Serialized
func GetCacheRenderViaSerialized(iSerialized int) (*CacheRender, error) {
	var _CacheRender = &CacheRender{Serialized: iSerialized}
	has, err := Engine.Get(_CacheRender)
	if has {
		return _CacheRender, err
	} else {
		return nil, err
	}
}

// GetCacheRenderViaTags Get CacheRender via Tags
func GetCacheRenderViaTags(iTags string) (*CacheRender, error) {
	var _CacheRender = &CacheRender{Tags: iTags}
	has, err := Engine.Get(_CacheRender)
	if has {
		return _CacheRender, err
	} else {
		return nil, err
	}
}

// GetCacheRenderViaChecksum Get CacheRender via Checksum
func GetCacheRenderViaChecksum(iChecksum string) (*CacheRender, error) {
	var _CacheRender = &CacheRender{Checksum: iChecksum}
	has, err := Engine.Get(_CacheRender)
	if has {
		return _CacheRender, err
	} else {
		return nil, err
	}
}

// SetCacheRenderViaCid Set CacheRender via Cid
func SetCacheRenderViaCid(iCid string, cache_render *CacheRender) (int64, error) {
	cache_render.Cid = iCid
	return Engine.Insert(cache_render)
}

// SetCacheRenderViaData Set CacheRender via Data
func SetCacheRenderViaData(iData []byte, cache_render *CacheRender) (int64, error) {
	cache_render.Data = iData
	return Engine.Insert(cache_render)
}

// SetCacheRenderViaExpire Set CacheRender via Expire
func SetCacheRenderViaExpire(iExpire int, cache_render *CacheRender) (int64, error) {
	cache_render.Expire = iExpire
	return Engine.Insert(cache_render)
}

// SetCacheRenderViaCreated Set CacheRender via Created
func SetCacheRenderViaCreated(iCreated string, cache_render *CacheRender) (int64, error) {
	cache_render.Created = iCreated
	return Engine.Insert(cache_render)
}

// SetCacheRenderViaSerialized Set CacheRender via Serialized
func SetCacheRenderViaSerialized(iSerialized int, cache_render *CacheRender) (int64, error) {
	cache_render.Serialized = iSerialized
	return Engine.Insert(cache_render)
}

// SetCacheRenderViaTags Set CacheRender via Tags
func SetCacheRenderViaTags(iTags string, cache_render *CacheRender) (int64, error) {
	cache_render.Tags = iTags
	return Engine.Insert(cache_render)
}

// SetCacheRenderViaChecksum Set CacheRender via Checksum
func SetCacheRenderViaChecksum(iChecksum string, cache_render *CacheRender) (int64, error) {
	cache_render.Checksum = iChecksum
	return Engine.Insert(cache_render)
}

// AddCacheRender Add CacheRender via all columns
func AddCacheRender(iCid string, iData []byte, iExpire int, iCreated string, iSerialized int, iTags string, iChecksum string) error {
	_CacheRender := &CacheRender{Cid: iCid, Data: iData, Expire: iExpire, Created: iCreated, Serialized: iSerialized, Tags: iTags, Checksum: iChecksum}
	if _, err := Engine.Insert(_CacheRender); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCacheRender Post CacheRender via iCacheRender
func PostCacheRender(iCacheRender *CacheRender) (string, error) {
	_, err := Engine.Insert(iCacheRender)
	return iCacheRender.Cid, err
}

// PutCacheRender Put CacheRender
func PutCacheRender(iCacheRender *CacheRender) (string, error) {
	_, err := Engine.Id(iCacheRender.Cid).Update(iCacheRender)
	return iCacheRender.Cid, err
}

// PutCacheRenderViaCid Put CacheRender via Cid
func PutCacheRenderViaCid(Cid_ string, iCacheRender *CacheRender) (int64, error) {
	row, err := Engine.Update(iCacheRender, &CacheRender{Cid: Cid_})
	return row, err
}

// PutCacheRenderViaData Put CacheRender via Data
func PutCacheRenderViaData(Data_ []byte, iCacheRender *CacheRender) (int64, error) {
	row, err := Engine.Update(iCacheRender, &CacheRender{Data: Data_})
	return row, err
}

// PutCacheRenderViaExpire Put CacheRender via Expire
func PutCacheRenderViaExpire(Expire_ int, iCacheRender *CacheRender) (int64, error) {
	row, err := Engine.Update(iCacheRender, &CacheRender{Expire: Expire_})
	return row, err
}

// PutCacheRenderViaCreated Put CacheRender via Created
func PutCacheRenderViaCreated(Created_ string, iCacheRender *CacheRender) (int64, error) {
	row, err := Engine.Update(iCacheRender, &CacheRender{Created: Created_})
	return row, err
}

// PutCacheRenderViaSerialized Put CacheRender via Serialized
func PutCacheRenderViaSerialized(Serialized_ int, iCacheRender *CacheRender) (int64, error) {
	row, err := Engine.Update(iCacheRender, &CacheRender{Serialized: Serialized_})
	return row, err
}

// PutCacheRenderViaTags Put CacheRender via Tags
func PutCacheRenderViaTags(Tags_ string, iCacheRender *CacheRender) (int64, error) {
	row, err := Engine.Update(iCacheRender, &CacheRender{Tags: Tags_})
	return row, err
}

// PutCacheRenderViaChecksum Put CacheRender via Checksum
func PutCacheRenderViaChecksum(Checksum_ string, iCacheRender *CacheRender) (int64, error) {
	row, err := Engine.Update(iCacheRender, &CacheRender{Checksum: Checksum_})
	return row, err
}

// UpdateCacheRenderViaCid via map[string]interface{}{}
func UpdateCacheRenderViaCid(iCid string, iCacheRenderMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheRender)).Where("cid = ?", iCid).Update(iCacheRenderMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheRenderViaData via map[string]interface{}{}
func UpdateCacheRenderViaData(iData []byte, iCacheRenderMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheRender)).Where("data = ?", iData).Update(iCacheRenderMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheRenderViaExpire via map[string]interface{}{}
func UpdateCacheRenderViaExpire(iExpire int, iCacheRenderMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheRender)).Where("expire = ?", iExpire).Update(iCacheRenderMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheRenderViaCreated via map[string]interface{}{}
func UpdateCacheRenderViaCreated(iCreated string, iCacheRenderMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheRender)).Where("created = ?", iCreated).Update(iCacheRenderMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheRenderViaSerialized via map[string]interface{}{}
func UpdateCacheRenderViaSerialized(iSerialized int, iCacheRenderMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheRender)).Where("serialized = ?", iSerialized).Update(iCacheRenderMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheRenderViaTags via map[string]interface{}{}
func UpdateCacheRenderViaTags(iTags string, iCacheRenderMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheRender)).Where("tags = ?", iTags).Update(iCacheRenderMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCacheRenderViaChecksum via map[string]interface{}{}
func UpdateCacheRenderViaChecksum(iChecksum string, iCacheRenderMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CacheRender)).Where("checksum = ?", iChecksum).Update(iCacheRenderMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCacheRender Delete CacheRender
func DeleteCacheRender(iCid string) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(CacheRender))
	return row, err
}

// DeleteCacheRenderViaCid Delete CacheRender via Cid
func DeleteCacheRenderViaCid(iCid string) (err error) {
	var has bool
	var _CacheRender = &CacheRender{Cid: iCid}
	if has, err = Engine.Get(_CacheRender); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CacheRender)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheRenderViaData Delete CacheRender via Data
func DeleteCacheRenderViaData(iData []byte) (err error) {
	var has bool
	var _CacheRender = &CacheRender{Data: iData}
	if has, err = Engine.Get(_CacheRender); (has == true) && (err == nil) {
		if row, err := Engine.Where("data = ?", iData).Delete(new(CacheRender)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheRenderViaExpire Delete CacheRender via Expire
func DeleteCacheRenderViaExpire(iExpire int) (err error) {
	var has bool
	var _CacheRender = &CacheRender{Expire: iExpire}
	if has, err = Engine.Get(_CacheRender); (has == true) && (err == nil) {
		if row, err := Engine.Where("expire = ?", iExpire).Delete(new(CacheRender)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheRenderViaCreated Delete CacheRender via Created
func DeleteCacheRenderViaCreated(iCreated string) (err error) {
	var has bool
	var _CacheRender = &CacheRender{Created: iCreated}
	if has, err = Engine.Get(_CacheRender); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(CacheRender)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheRenderViaSerialized Delete CacheRender via Serialized
func DeleteCacheRenderViaSerialized(iSerialized int) (err error) {
	var has bool
	var _CacheRender = &CacheRender{Serialized: iSerialized}
	if has, err = Engine.Get(_CacheRender); (has == true) && (err == nil) {
		if row, err := Engine.Where("serialized = ?", iSerialized).Delete(new(CacheRender)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheRenderViaTags Delete CacheRender via Tags
func DeleteCacheRenderViaTags(iTags string) (err error) {
	var has bool
	var _CacheRender = &CacheRender{Tags: iTags}
	if has, err = Engine.Get(_CacheRender); (has == true) && (err == nil) {
		if row, err := Engine.Where("tags = ?", iTags).Delete(new(CacheRender)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCacheRenderViaChecksum Delete CacheRender via Checksum
func DeleteCacheRenderViaChecksum(iChecksum string) (err error) {
	var has bool
	var _CacheRender = &CacheRender{Checksum: iChecksum}
	if has, err = Engine.Get(_CacheRender); (has == true) && (err == nil) {
		if row, err := Engine.Where("checksum = ?", iChecksum).Delete(new(CacheRender)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
