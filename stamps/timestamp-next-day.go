package stamps

func (stamp *Timestamp) NextDay() *Timestamp {
	ts := stamp.AddDate(0, 0, 1)
	return &Timestamp{&ts}
}
