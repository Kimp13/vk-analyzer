module main

go 1.15

replace local/logger => ../logger

require (
	github.com/fatih/color v1.10.0
	local/logger v0.0.0-00010101000000-000000000000
)
