# scripts for i3blocks

## time
Shows the time in different colors, depending on the current time (hour)

### Configuration
```
make i3b_time
```
in your i3blocks.conf:
```
# Date Time
[time]
command=/path/to/binary/i3blocks/i3b_time
interval=1
```

## battery
Thinkpads sometimes have multiple batteries
```
make i3b_battery
```
in your i3blocks.conf:
```
# Battery
[battery]
command=/path/to/binary/i3blocks/i3b_battery
interval=30
```