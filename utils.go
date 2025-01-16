package beidou


func Filter(tokens []Token,call func(token Token,index int) bool) []Token {
	_token := make([]Token,0)
	for index,token := range tokens {
		if(call(token,index)) {
			_token = append(_token, token)
		}
	}
	return _token
}