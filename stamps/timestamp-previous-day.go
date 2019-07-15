package stamps

func (stamp *Timestamp) PreviousDay() *Timestamp {
	ts := stamp.AddDate(0, 0, -1)
	return &Timestamp{&ts}
}
