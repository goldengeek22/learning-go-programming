package greeting

func salutation() string {
	if Polite {
		return "Good day"
	}
	return "Hey"
}
