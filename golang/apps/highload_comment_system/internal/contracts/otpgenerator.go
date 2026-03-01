package contracts

type OtpGenerator interface {
	Generate(length int) (string, error)
}
