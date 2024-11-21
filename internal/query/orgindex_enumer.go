// Code generated by "enumer -type orgIndex -linecomment"; DO NOT EDIT.

package query

import (
	"fmt"
	"strings"
)

const _orgIndexName = "orgIndexByIDorgIndexByPrimaryDomain"

var _orgIndexIndex = [...]uint8{0, 0, 12, 35}

const _orgIndexLowerName = "orgindexbyidorgindexbyprimarydomain"

func (i orgIndex) String() string {
	if i < 0 || i >= orgIndex(len(_orgIndexIndex)-1) {
		return fmt.Sprintf("orgIndex(%d)", i)
	}
	return _orgIndexName[_orgIndexIndex[i]:_orgIndexIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _orgIndexNoOp() {
	var x [1]struct{}
	_ = x[orgIndexUnspecified-(0)]
	_ = x[orgIndexByID-(1)]
	_ = x[orgIndexByPrimaryDomain-(2)]
}

var _orgIndexValues = []orgIndex{orgIndexUnspecified, orgIndexByID, orgIndexByPrimaryDomain}

var _orgIndexNameToValueMap = map[string]orgIndex{
	_orgIndexName[0:0]:        orgIndexUnspecified,
	_orgIndexLowerName[0:0]:   orgIndexUnspecified,
	_orgIndexName[0:12]:       orgIndexByID,
	_orgIndexLowerName[0:12]:  orgIndexByID,
	_orgIndexName[12:35]:      orgIndexByPrimaryDomain,
	_orgIndexLowerName[12:35]: orgIndexByPrimaryDomain,
}

var _orgIndexNames = []string{
	_orgIndexName[0:0],
	_orgIndexName[0:12],
	_orgIndexName[12:35],
}

// orgIndexString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func orgIndexString(s string) (orgIndex, error) {
	if val, ok := _orgIndexNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _orgIndexNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to orgIndex values", s)
}

// orgIndexValues returns all values of the enum
func orgIndexValues() []orgIndex {
	return _orgIndexValues
}

// orgIndexStrings returns a slice of all String values of the enum
func orgIndexStrings() []string {
	strs := make([]string, len(_orgIndexNames))
	copy(strs, _orgIndexNames)
	return strs
}

// IsAorgIndex returns "true" if the value is listed in the enum definition. "false" otherwise
func (i orgIndex) IsAorgIndex() bool {
	for _, v := range _orgIndexValues {
		if i == v {
			return true
		}
	}
	return false
}
