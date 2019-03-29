package marshal

type KeyPath struct {
	path []string
	loc int
}

func newKeyPath() *KeyPath {
	kp:=&KeyPath{}
	kp.path = make([]string,0)
	kp.loc = -1
	return kp
}

func (kp *KeyPath) add(key string) {
	kp.path = append(kp.path,key)
	kp.loc++
}

func (kp *KeyPath) set(key string) {
	kp.path[kp.loc]=key
}

func (kp *KeyPath) del(){
	kp.path = kp.path[0:kp.loc]
	kp.loc--
}