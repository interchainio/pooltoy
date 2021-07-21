package utils

import (
	"errors"
	"testing"
)

func Test_parseEmoji(t *testing.T) {
	emojis := []string{
		"\U0001f194",
		"1\U0001f194",
		"\U0001f9d1\u200d\U0001f680",
		"1\U0001f9d1\u200d\U0001f680",
	}

	for _, v := range emojis[:2] {
		emo, err := ParseEmoji(v)
		if err != nil {
			t.Fatal(err)
		}

		if emo != "\U0001f194" {
			t.Fatal(emo)
		}
	}

	for _, v := range emojis[2:] {
		emo, err := ParseEmoji(v)
		if err != nil {
			t.Fatal(err)
		}

		if emo != "\U0001f9d1\u200d\U0001f680" {
			t.Fatal(emo)
		}
	}

	emojis = []string{
		"100\U0001f194",
		"2\U0001f194",
		"1\U0001f194xyz",
		"0\U0001f194",
		"\U0001f194abc",
		"100\U0001f9d1\u200d\U0001f680",
		"2\U0001f9d1\u200d\U0001f680",
		"1\U0001f9d1\u200d\U0001f680xyz",
		"\U0001f9d1\u200d\U0001f680abc",
	}

	for _, v := range emojis {
		emo, err := ParseEmoji(v)

		if !errors.Is(err, ErrParseEmoji) {
			t.Fatal(err)
		}
		if emo != "" {
			t.Fatal(emo)
		}
	}
}
