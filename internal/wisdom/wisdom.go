package wisdom

import "errors"

var Wisdoms = map[int]string{
	1:  "The fool doth think he is wise, but the wise man knows himself to be a fool. ― William Shakespeare, As You Like It ",
	2:  "It is better to remain silent at the risk of being thought a fool, than to talk and remove all doubt of it. ― Maurice Switzer, Mrs. Goose, Her Book ",
	3:  "Whenever you find yourself on the side of the majority, it is time to reform (or pause and reflect). ― Mark Twain ",
	4:  "When someone loves you, the way they talk about you is different. You feel safe and comfortable. ― Jess C. Scott, The Intern ",
	5:  "Knowing yourself is the beginning of all wisdom. ― Aristotle ",
	6:  "The only true wisdom is in knowing you know nothing. ― Socrates ",
	7:  "The saddest aspect of life right now is that science gathers knowledge faster than society gathers wisdom. ― Isaac Asimov ",
	8:  "Hold fast to dreams, For if dreams die Life is a broken-winged bird, That cannot fly. ― Langston Hughes ",
	9:  "Count your age by friends, not years. Count your life by smiles, not tears. ― John Lennon ",
	10: "In a good bookroom you feel in some mysterious way that you are absorbing the wisdom contained in all the books through your skin, without even opening them. ― Mark Twain ",
	11: "May you live every day of your life. ― Jonathan Swift ",
	12: "Any fool can know. The point is to understand. ― Albert Einstein ",
	13: "It is the mark of an educated mind to be able to entertain a thought without accepting it. ― Aristotle, Metaphysics ",
	14: "The secret of life, though, is to fall seven times and to get up eight times. ― Paulo Coelho, The Alchemist ",
	15: "Never laugh at live dragons. ― J.R.R. Tolkien ",
	16: "If you're reading this... Congratulations, you're alive. If that's not something to smile about, then I don't know what is. ― Chad Sugg, Monsters Under Your Head ",
	17: "Think before you speak. Read before you think. ― Fran Lebowitz, The Fran Lebowitz Reader ",
	18: "Never let your sense of morals prevent you from doing what is right. ― Isaac Asimov, Foundation ",
	19: "The best index to a person's character is how he treats people who can't do him any good, and how he treats people who can't fight back. ― Abigail Van Buren ",
	20: "The unexamined life is not worth living. ― Socrates ",
}

func New() (string, error) {
	if len(Wisdoms) == 0 {
		return "", errors.New("wisdoms is empty")
	}

	wisdom := ""
	for _, wisdom = range Wisdoms {
		break
	}

	return wisdom, nil
}
