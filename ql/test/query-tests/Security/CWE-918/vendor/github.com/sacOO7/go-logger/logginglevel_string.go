// Code generated by "stringer -type=LoggingLevel"; DO NOT EDIT.

package logging

import "strconv"

const _LoggingLevel_name = "TRACEINFOWARNINGERROROFF"

var _LoggingLevel_index = [...]uint8{0, 5, 9, 16, 21, 24}

func (i LoggingLevel) String() string {
	if i < 0 || i >= LoggingLevel(len(_LoggingLevel_index)-1) {
		return "LoggingLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LoggingLevel_name[_LoggingLevel_index[i]:_LoggingLevel_index[i+1]]
}
