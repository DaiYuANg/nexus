package translate

import "embed"

//go:embed zh/*
var ZHI18n embed.FS

//go:embed en/*
var ENI18n embed.FS
