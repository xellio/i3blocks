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
command=/path/to/binary/i3blocks/i3b_time; bash -c 'if [ -n "$BLOCK_BUTTON" ]; then /path/to/binary/i3blocks/i3b_calendar; fi';
interval=1
```
in your i3 config:
```
for_window [class="Yad"] floating enable border pixel 0
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