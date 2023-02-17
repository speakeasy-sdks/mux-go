package shared

type SigningKey struct {
	CreatedAt  *string `json:"created_at,omitempty"`
	ID         *string `json:"id,omitempty"`
	PrivateKey *string `json:"private_key,omitempty"`
}
