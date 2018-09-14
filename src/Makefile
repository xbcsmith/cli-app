CC = $(shell which clang)
CFLAGS = -Wall -lm -I'.'
CFLAGS += $(shell pkg-config --cflags json-c)
LDFLAGS += $(shell pkg-config --libs json-c)

all: nodeinfo

nodeinfo: nodeinfo.c
	$(CC) $(CFLAGS) $(LDFLAGS) $(TARGET_ARCH) -o $@ $<

install:
	install -m 0755 -d $(DESTDIR)/bin
	install -m 0755 nodeinfo $(DESTDIR)/bin/nodeinfo
	rm -f nodeinfo

clean:
	rm -f nodeinfo

.PHONY:	all	install clean
