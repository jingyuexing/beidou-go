package beidou_test

import (
	"beidou"
	"testing"
)


func TestSOSMessage(t *testing.T) {
	tokens := beidou.Tokenizer("$BDEPI,x.x,x.x,x,hhmmss.ss,yyyyy.yy,c,llll.ll,c,x.x,x,x,x.x,c--c*hh\n\r")
	SOSMessage := beidou.Filter(tokens,func(token beidou.Token, index int) bool {
		return token.Type != beidou.Spiliter && token.Type != beidou.End && token.Type != beidou.Begin
	})
	if len(tokens) < 3  || len(SOSMessage) < 3{
		t.Error("Error:")
	}

	communicationMessage := beidou.Tokenizer("$BDTCI,x.x,x.x,hhmmss,x,x,x.x,c-c*hh\n\r")
	communicationMessageList := beidou.Filter(communicationMessage,func(token beidou.Token, index int) bool {
		return token.Type != beidou.Spiliter && token.Type != beidou.End && token.Type != beidou.Begin
	})
	if len(communicationMessageList) < 3 {
		t.Error("Error:")
	}

}