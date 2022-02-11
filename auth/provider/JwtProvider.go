package provider

type JwtProviderImpl struct {
}

func NewJwtProviderImpl() *JwtProviderImpl {
	return &JwtProviderImpl{}
}

func (j *JwtProviderImpl) encode() string {
	//TODO implement me
	panic("implement me")
}

func (j *JwtProviderImpl) decode() any {
	//TODO implement me
	panic("implement me")
}
