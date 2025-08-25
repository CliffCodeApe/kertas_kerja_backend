package token

type UserAuthToken struct {
	ID          uint64 `json:"id"`
	Nama_Satker string `json:"user_id"`
	KodeKL      string `json:"kode_kl"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Password    string `json:"password"`
}
