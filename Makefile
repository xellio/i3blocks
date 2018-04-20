PREFIX  = i3b
GO      = go

all: $(PREFIX)_time 

$(PREFIX)_time:
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_time ./time/main.go

$(PREFIX)_battery:
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_battery ./time/main.go

