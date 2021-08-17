package pagination

func GetParams(now, total, next, prev int, _category string) (category, nextParams, PrevParams interface{}) {
	if _category == "" {
		category = nil
		if now == 1 && next == 1 {
			nextParams = nil
			PrevParams = nil
			return
		} else if now == 1 {
			nextParams = next
			PrevParams = nil
			return
		} else if now == total {
			nextParams = nil
			PrevParams = prev
			return
		} else {
			nextParams = next
			PrevParams = prev
		}
	} else {
		category = _category
		if now == 1 && next == 1 {
			nextParams = nil
			PrevParams = nil
			return
		} else if now == 1 {
			nextParams = next
			PrevParams = nil
			return
		} else if now == total {
			nextParams = nil
			PrevParams = prev
			return
		} else {
			nextParams = next
			PrevParams = prev
		}
	}

	return
}

func GetNippoParams(now, total, next, prev int) (nextParams, PrevParams interface{}) {
	if now == 1 {
		nextParams = next
		PrevParams = nil
		return
	} else if now == total {
		nextParams = nil
		PrevParams = prev
		return
	}
	return next, prev
}
