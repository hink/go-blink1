package libusb

/*
	#cgo LDFLAGS: -lusb
	#include <usb.h>
*/
import "C"
import "unsafe"

const (
	USBRQ_HID_SET_REPORT        = 0x09
	USB_HID_REPORT_TYPE_FEATURE = 3
)

func SendBlink1Command(device *Device, fadeTime int, red, blue, green, led uint8) int {
	dms := fadeTime / 10

	data := []byte{
		'0', 'c', byte(red), byte(green), byte(blue), byte(dms >> 8), byte(dms % 127), byte(led),
	}

	//reportID := data[1]

	return int(C.usb_control_msg(device.handle,
		C.int(USB_TYPE_CLASS|C.USB_RECIP_INTERFACE|C.USB_ENDPOINT_OUT),
		C.int(USBRQ_HID_SET_REPORT),
		C.int(USB_HID_REPORT_TYPE_FEATURE<<8|('c'&0xff)),
		C.int(0),
		(*C.char)(unsafe.Pointer(&data[0])),
		C.int(len(data)),
		C.int(device.timeout)))
}
