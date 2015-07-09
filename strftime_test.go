package strftime

import (
	"bytes"
	"testing"
	"time"
)

type TestCase struct {
	format, value string
}

var testTime = time.Date(2009, time.November, 10, 23, 1, 2, 3, time.UTC)
var testCases = []*TestCase{
	&TestCase{"%a", "Tue"},
	&TestCase{"%A", "Tuesday"},
	&TestCase{"%b", "Nov"},
	&TestCase{"%B", "November"},
	&TestCase{"%c", "Tue, 10 Nov 2009 23:01:02 UTC"},
	&TestCase{"%d", "10"},
	&TestCase{"%H", "23"},
	&TestCase{"%I", "11"},
	&TestCase{"%j", "314"},
	&TestCase{"%m", "11"},
	&TestCase{"%M", "01"},
	&TestCase{"%p", "PM"},
	&TestCase{"%S", "02"},
	&TestCase{"%U", "45"},
	&TestCase{"%w", "2"},
	&TestCase{"%W", "45"},
	&TestCase{"%x", "11/10/09"},
	&TestCase{"%X", "23:01:02"},
	&TestCase{"%y", "09"},
	&TestCase{"%Y", "2009"},
	&TestCase{"%Z", "UTC"},
	&TestCase{"%3n", "000"},
	&TestCase{"%6n", "000000"},
	&TestCase{"%9n", "000000003"},

	// Escape
	&TestCase{"%%%Y", "%2009"},
	&TestCase{"%3%%", "%3%"},
	&TestCase{"%3%3n", "%3000"},
	&TestCase{"%3xy%3n", "%3xy000"},
	// Embedded
	&TestCase{"/path/%Y/%m/report", "/path/2009/11/report"},
	//Empty
	&TestCase{"", ""},
}

func TestFormats(t *testing.T) {
	for _, tc := range testCases {
		value := Format(tc.format, testTime)
		if value != tc.value {
			t.Fatalf("error in %s: got %s instead of %s", tc.format, value, tc.value)
		}
	}
}

func TestUnknown(t *testing.T) {
	unknownFormat := "%g"
	value := Format(unknownFormat, testTime)
	if unknownFormat != value {
		t.Fatalf("error to in %s: got %s instead of %s", unknownFormat, value, unknownFormat)
	}
}

func TestFormatter_ValidFormats(t *testing.T) {
	for _, tc := range testCases {
		formatter := NewFormatter(tc.format)
		value := formatter.Format(testTime)
		if value != tc.value {
			t.Fatalf("error in %s: got %s instead of %s", tc.format, value, tc.value)
		}
		buf := bytes.NewBuffer(make([]byte, 0, 0))
		formatter.FormatTo(buf, testTime)
		if string(buf.Bytes()) != tc.value {
			t.Fatalf("error in %s: got %s instead of %s", tc.format, value, tc.value)
		}
	}
}

func TestFormatter_InvalidFormats(t *testing.T) {
	unknownFormat := "%g"
	formatter := NewFormatter(unknownFormat)
	value := formatter.Format(testTime)
	if unknownFormat != value {
		t.Fatalf("error to in %s: get %s instead of %s", unknownFormat, value, unknownFormat)
	}
	buf := bytes.NewBuffer(make([]byte, 0, 0))
	formatter.FormatTo(buf, testTime)
	if unknownFormat != value {
		t.Fatalf("error to in %s: get %s instead of %s", unknownFormat, value, unknownFormat)
	}
}
