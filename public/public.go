package public

import "embed"

//go:embed static/*
var Static embed.FS

// //go:embed app/*
// var AppStatic embed.FS

// Tips👇
// https://www.cnblogs.com/apocelipes/p/13907858.html
