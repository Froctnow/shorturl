package url

const ServerURL = "http://localhost:8080"

func (u *urlUseCase) CreateShortURL(url string) string {
	urlEntity := u.provider.CreateURL(url)

	return ServerURL + "/" + urlEntity.ID
}
