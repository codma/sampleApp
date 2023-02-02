package model

type User struct {
	UserId   string  `json:"user_id" `   // 사용자 아이디
	UserName string  `json:"user_name" ` // 반려묘/견 이름
	Birthday string  `json:"birthday" `  // 생일
	Weight   float64 `json:"weight" `    //몸무게
	Type     string  `json:"type" `      //종류
}

type CreateUserResponse struct {
	MemberId uint64
}
