CC = $(shell which clang)
CFLAGS = -Wall -lm -I'.'
CFLAGS += $(shell pkg-config --cflags json-c)
LDFLAGS += $(shell pkg-config --libs json-c)

all: systeminfo

systeminfo: systeminfo.c
	$(CC) $(CFLAGS) $(LDFLAGS) $(TARGET_ARCH) -o $@ $<

install:
	install -m 0755 -d $(DESTDIR)/bin
	install -m 0755 systeminfo $(DESTDIR)/bin/systeminfo
	rm -f systeminfo

clean:
	rm -f systeminfo

.PHONY:	all	install clean
