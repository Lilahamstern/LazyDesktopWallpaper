package win

import (
	"github.com/lilahamstern/automaticDesktopWallpaper/pkg/file"
	"golang.org/x/sys/windows"
	"unsafe"
)

var (
	user32DLL           = windows.NewLazyDLL("user32.dll")
	procSystemParamInfo = user32DLL.NewProc("SystemParametersInfoW")
)

func ChangeDesktopBackground() {
	imagePath, _ := windows.UTF16PtrFromString(file.GetImageFilePath())
	procSystemParamInfo.Call(20, 0, uintptr(unsafe.Pointer(imagePath)), 0x001A)
}
