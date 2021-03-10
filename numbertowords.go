package numbertowords

import (
	"errors"
)

var words = [21]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tenwords = [10]string{
	"",
	"",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

//maxnumber Maximum number should be valid
const maxnumber = 99999

//Convert converts number between 0 to Maxnumber to words.
func Convert(number int) (string, error) {
	if number < 0 || number > maxnumber {
		return "", errors.New("Number is not in valid range")
	}

	result := ""
	thousundresult := ""

	thousund := number / 1000
	if thousund > 0 {
		if thousund <= 19 {
			thousundresult = words[thousund] + " thousund "
		} else if thousund > 19 && thousund <= 99 {
			temptens := thousund / 10
			tempunits := thousund % 10
			thousundresult = tenwords[temptens] + words[tempunits] + " thousund "
		}
		number = number % 1000
		if number == 0 {
			return thousundresult, nil
		}
	}

	hundreds := number / 100

	if hundreds > 0 {
		result = words[hundreds] + " hundred "
		number = number % 100
		if number == 0 {
			return result, nil
		}
	}

	tens := number / 10
	units := number % 10

	if tens < 2 {
		return thousundresult + result + words[number], nil
	}

	if units == 0 {
		return thousundresult + result + tenwords[tens], nil
	}

	return thousundresult + result + tenwords[tens] + " " + words[units], nil

}
