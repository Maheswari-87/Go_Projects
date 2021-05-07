package greeting

import "testing"

func TestHello(t *testing.T) {
	emptyResult := HelloM("")
	if emptyResult != "Hello dude!" {
		t.Errorf("hello(\"\") Failed, Expected %v, got %v", "Hello dude!", emptyResult)
	} else {
		t.Logf("hello(\"\") Success, Expected %v, got %v", "Hello dude!", emptyResult)
	}
}
func TestWithoutPunctuation(t *testing.T) {
	result := HelloM("Mahi")
	if result != "Hello Mahi!" {
		t.Errorf("hello(\"Mahi\") Failed, Expected %v, got %v", "Hello Mahi!", result)
	} else {
		t.Logf("hello(\"Mahi\") Success, Expected %v, got %v", "Hello Mahi!", result)
	}
}
