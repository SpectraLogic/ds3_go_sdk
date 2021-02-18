package networking

// Used to correctly sort headers when creating the stringToSign for the authorization header.
type CanonicalHeader struct{
	key string // key is assumed to be lower cased
	values []string
}
type CanonicalHeaders []CanonicalHeader

func (p CanonicalHeaders) Len() int           { return len(p) }
func (p CanonicalHeaders) Less(i, j int) bool { return p[i].key < p[j].key }
func (p CanonicalHeaders) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
