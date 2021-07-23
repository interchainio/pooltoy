package utils

// parse emoji for mintfor, only allows format like "\U0001f630" or "1\U0001f630"
func ParseEmoji(emoji string) (string, error) {
	emojiMap := ReverseMapKV(EmojiCodeMap)
	_, found1 := emojiMap[emoji]
	_, found2 := emojiMap[emoji[1:]]

	if found1{
		return emoji, nil
	}

	if emoji[0] == '1' && found2 {
		return emoji[1:], nil
	}

	return "",ErrParseEmoji
}


func ReverseMapKV(emojiMap map[string]string)map[string]string{
	reversedMap := map[string]string{}
	for k, v := range emojiMap{
		reversedMap[v] = k
	}

	return reversedMap
}
