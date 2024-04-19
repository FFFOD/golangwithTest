package main

import "testing"

func TestHello(t *testing.T) {
	// 辅助函数
	assertCorrectMessage := func(t *testing.T, got, want string) {
		//t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
