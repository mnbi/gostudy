package main

import (
	"testing"
)

func assertBool(t *testing.T, msg string, result bool, expected bool) {
	if result != expected {
		t.Fatalf(msg, result, expected)
	}
}

func assertRune(t *testing.T, msg string, result rune, expected rune) {
	if result != expected {
		t.Fatalf(msg, result, expected)
	}
}

func TestIsHirakana(t *testing.T) {
	testRune := 'あ'
	expected := true
	result := isHirakana(testRune)
	assertBool(t, `isHirakana() = %v, want %v`, result, expected)
}

func TestIsKatakana(t *testing.T) {
	testRune := 'ア'
	expected := true
	result := isKatakana(testRune)
	assertBool(t, `isKatakana() = %v, want %v`, result, expected)
}

func TestToHirakana(t *testing.T) {
	testRune := 'イ'
	expected := 'い'
	result := toHirakana(testRune)
	assertRune(t, `toHirakana() = %c, want %c`, result, expected)
}

func TestToKatakana(t *testing.T) {
	testRune := 'う'
	expected := 'ウ'
	result := toKatakana(testRune)
	assertRune(t, `toKatakana() = %c, want %c`, result, expected)
}

func TestToHirakanaDakuon(t *testing.T) {
	testRune := 'ブ'
	expected := 'ぶ'
	result := toHirakana(testRune)
	assertRune(t, `toHirakana() = %c, want %c`, result, expected)
}

func TestToKatakanaDakuon(t *testing.T) {
	testRune := 'ぢ'
	expected := 'ヂ'
	result := toKatakana(testRune)
	assertRune(t, `toKatakana() = %c, want %c`, result, expected)
}

func TestToHirakanaHandakuon(t *testing.T) {
	testRune := 'パ'
	expected := 'ぱ'
	result := toHirakana(testRune)
	assertRune(t, `toHirakana() = %c, want %c`, result, expected)
}

func TestToKatakanaHandakuon(t *testing.T) {
	testRune := 'ぽ'
	expected := 'ポ'
	result := toKatakana(testRune)
	assertRune(t, `toKatakana() = %c, want %c`, result, expected)
}

func TestTrString(t *testing.T) {
	testStr := "apple ORANGE"
	expected := "APPLE orange"
	result := tr(testStr)
	if result != expected {
		t.Fatalf(`tr() = "%v", want "%v"`, result, expected)
	}
}
