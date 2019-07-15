package notifier

func new() *Notifier {
	return &Notifier{
		Topics: make(map[string]*Topic),
	}
}
