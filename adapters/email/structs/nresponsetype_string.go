// Code generated by "stringer -type=NResponseType"; DO NOT EDIT

package structs

import "fmt"

const _NResponseType_name = "NRspTUnknownNRspTServicesNRspTServicesAreaNRspTCreateNRspTSearchLLNRspTSearchDID"

var _NResponseType_index = [...]uint8{0, 12, 25, 42, 53, 66, 80}

func (i NResponseType) String() string {
	if i < 0 || i >= NResponseType(len(_NResponseType_index)-1) {
		return fmt.Sprintf("NResponseType(%d)", i)
	}
	return _NResponseType_name[_NResponseType_index[i]:_NResponseType_index[i+1]]
}
