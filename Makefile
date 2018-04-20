PREFIX  = i3b
GO      = go

all: clean $(PREFIX)_time $(PREFIX)_battery

$(PREFIX)_time:
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_time ./time/main.go

$(PREFIX)_battery:
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_battery ./battery/main.go

clean:
	rm -f $(PREFIX)_time $(PREFIX)_battery

.PHONY:all