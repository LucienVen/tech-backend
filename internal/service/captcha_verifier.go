package service

type CaptchaVerifier interface {
	Verify(id, code string, clear bool) bool
	Generate() (id string, b64s string, err error)
}
