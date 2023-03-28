package main

// Implemented with reference to "A Regular Expression Matcher" by Rob Pike

//    /* match: search for regexp anywhere in text */
//    int match(char *regexp, char *text)
//    {
//        if (regexp[0] == '^')
//            return matchhere(regexp+1, text);
//        do {    /* must look even if string is empty */
//            if (matchhere(regexp, text))
//                return 1;
//        } while (*text++ != '\0');
//        return 0;
//    }

// Match search for regexp anywhere in text
func Match(regexp, text string) bool {
	if regexp != "" && regexp[0] == '^' {
		return matchHere(regexp[1:], text)
	}
	for {
		if matchHere(regexp, text) {
			return true
		}
		if text == "" {
			return false
		}
		text = text[1:]
	}
}

//    /* matchhere: search for regexp at beginning of text */
//    int matchhere(char *regexp, char *text)
//    {
//        if (regexp[0] == '\0')
//            return 1;
//        if (regexp[1] == '*')
//            return matchstar(regexp[0], regexp+2, text);
//        if (regexp[0] == '$' && regexp[1] == '\0')
//            return *text == '\0';
//        if (*text!='\0' && (regexp[0]=='.' || regexp[0]==*text))
//            return matchhere(regexp+1, text+1);
//        return 0;
//    }

// matchHere search for regexp at beginning of text
func matchHere(regexp, text string) bool {
	switch {
	case regexp == "":
		return true
	case regexp == "$":
		return text == ""
	case len(regexp) >= 2 && regexp[1] == '*':
		return matchStar(regexp[0], regexp[2:], text)
	case text != "" && (regexp[0] == '.' || regexp[0] == text[0]):
		return matchHere(regexp[1:], text[1:])
	}
	return false
}

//    /* matchstar: search for c*regexp at beginning of text */
//    int matchstar(int c, char *regexp, char *text)
//    {
//        do {    /* a * matches zero or more instances */
//            if (matchhere(regexp, text))
//                return 1;
//        } while (*text != '\0' && (*text++ == c || c == '.'));
//        return 0;
//    }

// matchStar search for c*regexp at beginning of text
func matchStar(c byte, regexp, text string) bool {
	for {
		if matchHere(regexp, text) {
			return true
		}
		if text == "" || (text[0] != c && c != '.') {
			return false
		}
		text = text[1:]
	}
}