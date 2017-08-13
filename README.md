# GoBannerNabber
Port Scanner written in Golang
It's a variation of the one I wrote in Python: https://github.com/continuousqa/bannernabber
# Usage
User runs the app, answers three questions (host to scan, starting port and ending port) and the app scans each one reporting any banner it gets from each port.

Note: Only use this tool on sites you have permission to test and always know the legalities of using a port scanner in your territory.

# Performance
In the current implementation, reading buffers as raw bytes (instead of using the bufio library) I could scan
8000 ports in about 20 seconds (previously bufio reads were taking over a min to scan 8000 ports.)

Compared to the Python script, the Go app is about 2x faster - I get about 40s for the python script to scan 8000 ports.

