package document

func fix(an Annotations) Annotations {
	for _, v := range an.Spans {
		if isPunct(string(an.Text[v.Start])) {
			v.Start++
		}
		if !isPunct(string(an.Text[v.End])) {
			if isSpace(string(an.Text[v.End])) {
				v.End--
			} else {
				v.End++
			}
		}
	}
}
