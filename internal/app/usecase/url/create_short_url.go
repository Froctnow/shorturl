package url

func (u *urlUseCase) CreateShortURL(url string) string {
	urlEntity := u.provider.CreateURL(url)

	return u.serverURL + "/" + urlEntity.ID
}
