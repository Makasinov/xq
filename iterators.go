package main

const (
	AL    = 60
	AR    = 62
	SLASH = 116
)

func al(s int32) bool {
	if s == AL {
		return true
	}
	return false
}

func ar(s int32) bool {
	if s == AR {
		return true
	}
	return false
}

func slash(s int32) bool {
	if s == SLASH {
		return true
	}
	return false
}
