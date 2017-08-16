# GoBannerNabber
Port Scanner written in Golang
It's a variation of the one I wrote in Python: https://github.com/continuousqa/bannernabber
# Usage
User runs the app, answers three questions (host to scan, starting port and ending port) and the app scans each one reporting any banner it gets from each port.

Note: Only use this tool on sites you have permission to test and always know the legalities of using a port scanner in your territory.

# Performance
I modified the code a bit to break out a port range into 10 concurrent Go Routines.  This concurrency
has allowed me to scan a 65,000 port range in about 1min.   I'm not sure where the Go Routine sweet spot would be.
The number of concurrent Go Routines could be bumped up further (as needed.)