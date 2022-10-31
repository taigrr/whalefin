# If PREFIX isn't provided, we check for /usr/local and use that if it exists.
# Otherwice we fall back to using /usr.

LOCAL != test -d $(DESTDIR)/usr/local && echo -n "/local" || echo -n ""
LOCAL ?= $(shell test -d $(DESTDIR)/usr/local && echo "/local" || echo "")
PREFIX ?= /usr$(LOCAL)

.PHONY: build

build:
	wails build 

install: build
	install -Dm00755 build/whalefin $(DESTDIR)$(PREFIX)/bin/whalefin
	install -Dm00644 whalefin.service /etc/systemd/system/whalefin.service
	test ! -f /etc/pam.d/display_manager && install -Dm00644 /etc/pam.d/login /etc/pam.d/display_manager || true

uninstall:
	-rm $(DESTDIR)$(PREFIX)/bin/whalefin
	-rm $(DESTDIR)$(PREFIX)/lib/systemd/system/whalefin.service

embed:
	DISPLAY=:0 Xephyr :1 -screen 1280x720 &
	DISPLAY=:1 wails serve
