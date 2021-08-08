package view

import "embed"

//go:embed admin/*
var Admin embed.FS

//go:embed static/*
var Static embed.FS

// Tips
// https://www.cnblogs.com/apocelipes/p/13907858.html
