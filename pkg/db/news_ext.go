package db

func (s *NewsSearch) WithTagID(tagID *int) *NewsSearch {
	s.With(`? = ANY (t."tagIds")`, *tagID)
	return s
}
