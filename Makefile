PREFIX  = i3b
GO      = go

GLIDE_VERSION := $(shell glide --version 2>/dev/null)
DEP_VERSION := $(shell dep version 2>/dev/null)

all: clean $(PREFIX)_time $(PREFIX)_calendar $(PREFIX)_battery

$(PREFIX)_time: 
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_time ./time/main.go

$(PREFIX)_calendar: vendor
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_calendar ./calendar/main.go

$(PREFIX)_battery:
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_battery ./battery/main.go

clean:
	rm -f $(PREFIX)_time $(PREFIX)_calendar $(PREFIX)_battery

vendor:
ifdef DEP_VERSION
	dep ensure
else ifdef GLIDE_VERSION
	glide install
else
	go get .
endif

.PHONY:all