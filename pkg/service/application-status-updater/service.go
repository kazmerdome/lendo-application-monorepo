package applicationstatusupdater

type Service interface {
	ConsumeFromApplicationProcessedQueue()
}
