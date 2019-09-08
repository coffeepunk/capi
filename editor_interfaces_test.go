package capi

import (
	"fmt"
	"testing"
)

func TestEditorInterfaceControl_Get(t *testing.T) {
	editorInterface := NewEditorInterfaceControls(cma)
	ei := editorInterface.Get("")

	fmt.Println(ei)
}

func TestEditorInterfaceControl_Update(t *testing.T) {
	editorInterface := NewEditorInterfaceControls(cma)
	editorInterface.Update("")
}
