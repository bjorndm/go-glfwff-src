//go:build freebsd || netbsd || openbsd
// +build freebsd netbsd openbsd

package glfw

/*
#ifdef _GLFW_WAYLAND
	#include "wl_init.c.i"
	#include "wl_monitor.c.i"
	#include "wl_window.c.i"
	#include "wayland-idle-inhibit-unstable-v1-client-protocol.c.i"
	#include "wayland-pointer-constraints-unstable-v1-client-protocol.c.i"
	#include "wayland-relative-pointer-unstable-v1-client-protocol.c.i"
	#include "wayland-viewporter-client-protocol.c.i"
	#include "wayland-xdg-decoration-unstable-v1-client-protocol.c.i"
	#include "wayland-xdg-shell-client-protocol.c.i"
#endif
#ifdef _GLFW_X11
	#include "x11_init.c.i"
	#include "x11_monitor.c.i"
	#include "x11_window.c.i"
	#include "glx_context.c.i"
#endif
#include "null_joystick.c.i"
#include "posix_module.c.i"
#include "posix_poll.c.i"
#include "posix_time.c.i"
#include "posix_thread.c.i"
#include "xkb_unicode.c.i"
#include "egl_context.c.i"


*/
import "C"
