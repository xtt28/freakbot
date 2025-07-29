package classifier

type ClassifierService interface {
	// IsFlagged returns whether a specific message has been flagged for
	// containing inappropriate content.
	IsFlagged(string) (bool, error)
}
