package main

import "embed"

//go:embed migrations
var MigrationFs embed.FS
