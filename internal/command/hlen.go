package command

// Returns the number of fields in a hash.
// HLEN key
// https://redis.io/commands/hlen
type HLen struct {
	baseCmd
	key string
}

func parseHLen(b baseCmd) (*HLen, error) {
	cmd := &HLen{baseCmd: b}
	if len(cmd.args) != 1 {
		return cmd, ErrInvalidArgNum(cmd.name)
	}
	cmd.key = string(cmd.args[0])
	return cmd, nil
}

func (cmd *HLen) Run(w Writer, red Redka) (any, error) {
	count, err := red.Hash().Len(cmd.key)
	if err != nil {
		w.WriteError(translateError(err))
		return nil, err
	}
	w.WriteInt(count)
	return count, nil
}
