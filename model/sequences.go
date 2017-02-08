package model

type Sequences struct {
	Value int64 `xorm:"not null pk autoincr INT(10)"`
}

// GetSequencesesCount Sequencess' Count
func GetSequencesesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Sequences{})
	return total, err
}

// GetSequencesCountViaValue Get Sequences via Value
func GetSequencesCountViaValue(iValue int64) int64 {
	n, _ := Engine.Where("value = ?", iValue).Count(&Sequences{Value: iValue})
	return n
}

// GetSequencesesByValue Get Sequenceses via Value
func GetSequencesesByValue(offset int, limit int, Value_ int64) (*[]*Sequences, error) {
	var _Sequences = new([]*Sequences)
	err := Engine.Table("sequences").Where("value = ?", Value_).Limit(limit, offset).Find(_Sequences)
	return _Sequences, err
}

// GetSequenceses Get Sequenceses via field
func GetSequenceses(offset int, limit int, field string) (*[]*Sequences, error) {
	var _Sequences = new([]*Sequences)
	err := Engine.Table("sequences").Limit(limit, offset).Desc(field).Find(_Sequences)
	return _Sequences, err
}

// GetSequencesesViaValue Get Sequencess via Value
func GetSequencesesViaValue(offset int, limit int, Value_ int64, field string) (*[]*Sequences, error) {
	var _Sequences = new([]*Sequences)
	err := Engine.Table("sequences").Where("value = ?", Value_).Limit(limit, offset).Desc(field).Find(_Sequences)
	return _Sequences, err
}

// HasSequencesViaValue Has Sequences via Value
func HasSequencesViaValue(iValue int64) bool {
	if has, err := Engine.Where("value = ?", iValue).Get(new(Sequences)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetSequencesViaValue Get Sequences via Value
func GetSequencesViaValue(iValue int64) (*Sequences, error) {
	var _Sequences = &Sequences{Value: iValue}
	has, err := Engine.Get(_Sequences)
	if has {
		return _Sequences, err
	} else {
		return nil, err
	}
}

// SetSequencesViaValue Set Sequences via Value
func SetSequencesViaValue(iValue int64, sequences *Sequences) (int64, error) {
	sequences.Value = iValue
	return Engine.Insert(sequences)
}

// AddSequences Add Sequences via all columns
func AddSequences(iValue int64) error {
	_Sequences := &Sequences{Value: iValue}
	if _, err := Engine.Insert(_Sequences); err != nil {
		return err
	} else {
		return nil
	}
}

// PostSequences Post Sequences via iSequences
func PostSequences(iSequences *Sequences) (int64, error) {
	_, err := Engine.Insert(iSequences)
	return iSequences.Value, err
}

// PutSequences Put Sequences
func PutSequences(iSequences *Sequences) (int64, error) {
	_, err := Engine.Id(iSequences.Value).Update(iSequences)
	return iSequences.Value, err
}

// PutSequencesViaValue Put Sequences via Value
func PutSequencesViaValue(Value_ int64, iSequences *Sequences) (int64, error) {
	row, err := Engine.Update(iSequences, &Sequences{Value: Value_})
	return row, err
}

// UpdateSequencesViaValue via map[string]interface{}{}
func UpdateSequencesViaValue(iValue int64, iSequencesMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Sequences)).Where("value = ?", iValue).Update(iSequencesMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteSequences Delete Sequences
func DeleteSequences(iValue int64) (int64, error) {
	row, err := Engine.Id(iValue).Delete(new(Sequences))
	return row, err
}

// DeleteSequencesViaValue Delete Sequences via Value
func DeleteSequencesViaValue(iValue int64) (err error) {
	var has bool
	var _Sequences = &Sequences{Value: iValue}
	if has, err = Engine.Get(_Sequences); (has == true) && (err == nil) {
		if row, err := Engine.Where("value = ?", iValue).Delete(new(Sequences)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
