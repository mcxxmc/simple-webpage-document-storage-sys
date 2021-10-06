package logging

import "go.uber.org/zap"

// SS string : string
type SS struct {
	S1 string
	S2 string
}

// converts *SS to zap.Field
func (ss *SS) convert() zap.Field {
	return zap.String(ss.S1, ss.S2)
}
