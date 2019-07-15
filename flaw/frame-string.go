package flaw

import "fmt"

func (frm *frame) String() string {
	return fmt.Sprintf("%s:%d", frm.Pathname, frm.Line)
}
