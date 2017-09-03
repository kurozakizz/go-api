package calculator

import "testing"

func Test_Add_EmptyString_ShouldReturn_Zero(t *testing.T) {
	if Add("") != 0 {
		t.Error()
	}
}

func Test_Add_1_ShouldReturn_1(t *testing.T) {
	if Add("1") != 1 {
		t.Error()
	}
}

func Test_Add_12_ShouldReturn_3(t *testing.T) {
	if Add("1,2") != 3 {
		t.Error()
	}
}

func Test_Add_1n2_ShouldReturn_3(t *testing.T) {
	if Add("1\n2") != 3 {
		t.Error()
	}
}

func Test_Add_1n23_ShoulReturn_6(t *testing.T) {
	if Add("1\n2,3") != 6 {
		t.Error()
	}
}

func Test_Add_1ncomma23_ShoulReturn_ErrorMessage(t *testing.T) {
	defer func() {
		if r := recover(); r != "Not need to prove it, just clarifying" {
			t.Error()
		}
	}()
	Add("1\n,2,3")
}
