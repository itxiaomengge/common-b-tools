package mlog

// Assert error
func Assert(err error) {
	if err != nil {
		panic(err)
	}
}
