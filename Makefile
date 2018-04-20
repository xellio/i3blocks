PREFIX  = i3b
GO      = go

all: $(PREFIX)_time 

$(PREFIX)_time:
	$(GO) build -ldflags="-s -w" -o $(PREFIX)_time ./time/main.go
