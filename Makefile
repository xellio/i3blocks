PREFIX  = i3b
GO      = go

GLIDE_VERSION := $(shell glide --version 2>/dev/null)
DEP_VERSION := $(shell dep version 2>/dev/null)

YAD_VERSION := $(shell yad --version 2>/dev/null)

all: clean \
	$(PREFIX)_time \
	$(PREFIX)_calendar \
	$(PREFIX)_battery \
	$(PREFIX)_battery_information

$(PREFIX)_time: 
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_time ./time/main.go

$(PREFIX)_calendar: dep_check
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_calendar ./calendar/main.go

$(PREFIX)_battery: vendor
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_battery ./battery/main.go

$(PREFIX)_battery_information: dep_check
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_battery_information ./battery_information/main.go

dep_check:
ifndef YAD_VERSION
	$(error "yad (Yet another dialoging program) is not available please install yad")
endif

clean:
	rm -rf vendor
	rm -f \
		$(PREFIX)_time \
		$(PREFIX)_calendar \
		$(PREFIX)_battery \
		$(PREFIX)_battery_information

vendor:
ifdef DEP_VERSION
	dep ensure
else ifdef GLIDE_VERSION
	glide install
else
	go get .
endif

.PHONY:all dep_check