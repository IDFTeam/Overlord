//go:build darwin && cgo

package sysinfo

/*
#cgo LDFLAGS: -framework ApplicationServices -framework CoreGraphics -framework CoreFoundation

#include <ApplicationServices/ApplicationServices.h>
#include <CoreFoundation/CoreFoundation.h>
#include <CoreGraphics/CoreGraphics.h>
#include <stdbool.h>

extern bool CGPreflightScreenCaptureAccess(void) __attribute__((weak_import));
extern bool CGRequestScreenCaptureAccess(void) __attribute__((weak_import));

static int checkAccessibilityPermission(void) {
	return AXIsProcessTrusted() ? 1 : 0;
}

static int checkScreenRecordingPermission(void) {
	if (CGPreflightScreenCaptureAccess == NULL) return 0;
	return CGPreflightScreenCaptureAccess() ? 1 : 0;
}

static int requestAccessibilityPermission(void) {
	const void *keys[] = { kAXTrustedCheckOptionPrompt };
	const void *values[] = { kCFBooleanTrue };
	CFDictionaryRef opts = CFDictionaryCreate(
		kCFAllocatorDefault,
		keys,
		values,
		1,
		&kCFCopyStringDictionaryKeyCallBacks,
		&kCFTypeDictionaryValueCallBacks
	);
	Boolean trusted = AXIsProcessTrustedWithOptions(opts);
	if (opts != NULL) CFRelease(opts);
	return trusted ? 1 : 0;
}

static int requestScreenRecordingPermission(void) {
	if (CGRequestScreenCaptureAccess == NULL) return checkScreenRecordingPermission();
	return CGRequestScreenCaptureAccess() ? 1 : 0;
}
*/
import "C"

func darwinAccessibilityPermission() bool {
	return C.checkAccessibilityPermission() == 1
}

func darwinScreenRecordingPermission() bool {
	return C.checkScreenRecordingPermission() == 1
}

func darwinRequestAccessibilityPermission() bool {
	return C.requestAccessibilityPermission() == 1
}

func darwinRequestScreenRecordingPermission() bool {
	return C.requestScreenRecordingPermission() == 1
}
