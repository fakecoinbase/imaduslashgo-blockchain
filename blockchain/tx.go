package blockchain

//TxInput struct
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

//TxOutput struct
type TxOutput struct {
	Value  int
	PubKey string
}

//CanUnlock checks if the user can unlock the transaction
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

//CanBeUnlocked checks if the user can unlock the transaction via the pubkey
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
