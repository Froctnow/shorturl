package url

func (u *urlUseCase) CreateShortURL(url string) (string, error) {
	urlEntity, err := u.provider.CreateURL(url)

	if err != nil {
		return "", err
	}

	return u.serverURL + "/" + urlEntity.ID, nil
}
