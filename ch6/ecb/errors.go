package ecb

type ecbankError string

func (e ecbankError) Error() string {
	return string(e)
}
