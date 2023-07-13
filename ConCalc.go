package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidOperator(testStrKey string) bool {

	type void struct{} //  sets helpers
	var voidEl void    //	sets helpers
	OpsSet := make(map[string]void)

	OpsSet["+"] = voidEl
	OpsSet["-"] = voidEl
	OpsSet["/"] = voidEl
	OpsSet["*"] = voidEl

	if _, ok := OpsSet[testStrKey]; ok {
		return true
	} else {
		return false
	}
}

func checkOperandsArab(testStrKey string) bool {

	type void struct{} //  sets helpers
	var voidEl void    //	sets helpers
	ArabSet := make(map[string]void)

	///ArabSet["0"] = voidEl
	ArabSet["1"] = voidEl
	ArabSet["2"] = voidEl
	ArabSet["3"] = voidEl
	ArabSet["4"] = voidEl
	ArabSet["5"] = voidEl
	ArabSet["6"] = voidEl
	ArabSet["7"] = voidEl
	ArabSet["8"] = voidEl
	ArabSet["9"] = voidEl
	ArabSet["10"] = voidEl
	if _, ok := ArabSet[testStrKey]; ok {
		return true
	} else {
		return false
	}
}

func checkOperandsRome(testStrKey string) bool {

	type void struct{} //  sets helpers
	var voidEl void    //	sets helpers
	RomeSet := make(map[string]void)

	RomeSet["I"] = voidEl
	RomeSet["II"] = voidEl
	RomeSet["III"] = voidEl
	RomeSet["IV"] = voidEl
	RomeSet["V"] = voidEl
	RomeSet["VI"] = voidEl
	RomeSet["VII"] = voidEl
	RomeSet["VIII"] = voidEl
	RomeSet["IX"] = voidEl
	RomeSet["X"] = voidEl

	if _, ok := RomeSet[testStrKey]; ok {
		return true
	} else {
		return false
	}
}

func intToRome100(i int) string {

	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman string
	for _, conversion := range conversions {
		for i >= conversion.value {
			roman = roman + conversion.digit
			i -= conversion.value
		}
	}

	return roman
}

const BasicErrorMsg = "То что Вы ввели перед тем как нажали \"Enter\" НЕ соответствует формату:\n{первый операнд арифметический оператор { + - * / } второй операнд}!"
const BasicGreetingMsg = "Приложение \"калькулятор\" снова выведет в консоль результат, если Вы снова введете:\n{число1 затем арифметический знак затем число2 и затем нажмете \"Enter\"} с пробелами или без"

func main() {
	RomeMap := make(map[string]int)

	RomeMap["I"] = 1
	RomeMap["II"] = 2
	RomeMap["III"] = 3
	RomeMap["IV"] = 4
	RomeMap["V"] = 5
	RomeMap["VI"] = 6
	RomeMap["VII"] = 7
	RomeMap["VIII"] = 8
	RomeMap["IX"] = 9
	RomeMap["X"] = 10

	var str1, str2, str3, str4, InputString, errorString, greetingString string
	var op1, op2, res int

	//проверяем не ввел ли пользователь четыре слова или более, заготовим строку, которую случайно никто не введет
	//а ее значение меняется на значение четвертого слова, вводимого пользователем
	str4 = "cheatcode543210string" //uniq string life user without special knowledge never enters
	errorString = BasicErrorMsg
	greetingString = "Приложение \"калькулятор\" выведет в консоль результат, если Вы введете:\n{число1 затем арифметический знак затем число2 затем нажмете \"Enter\"} с пробелами или без"

	//main cycle
	for {

		fmt.Println(greetingString)       //приветсвенный коментарий, впервые запускается без "опять"
		greetingString = BasicGreetingMsg //а со второй итерацпии уже с "опять"

		fmt.Scanln(&str1, &str2, &str3, &str4)

		InputString += str1 //если пользователь таки использовал пробелы сначала собираем все в одну строку
		InputString += str2 //но разобьем ее снова на три два операнда и знак в следующем блоке
		InputString += str3 //который совмещен с проверкой на наличие арифметического знака

		//проверка c подсказкой пользователю что пользователь ввел мало
		if len(InputString) <= 2 {
			errorString = ("Учтите, Вы ввели или слишком мало, вы же в курсе..? , что \n" + BasicErrorMsg)
			break
		}
		//изменено ли четвертое слово вводом пользователя? А четыре слова это ошибочно много
		if str4 != "cheatcode543210string" {
			errorString = ("Учтите, Вы ввели более трех слов, т.е.\n" + BasicErrorMsg)
			break
		}
		// там где пользователю можно без проблелов. Проверка на валидный оператор будет другая
		// c целью минимизировать неректность в советах пользователю вводить правильно проверям не ввел ли пользователь знак в конце, и упрекаем пользоваля именно в этом
		/* самое сложное место скрипта, мы тут вылавливаем ситуации когда пользоватеь вводит "3 3 +" или "7 7+" */

		if isValidOperator(str3) || (str2 != "" && !isValidOperator(str2) && isValidOperator(str2[len(str2)-1:])) {
			errorString = ("Учтите, Вы как мнимум ввели арифметический знак из ряда + - * / В конце, а надо в середине!\nИ вполне может быть что-то еще не то ввели\n" + BasicErrorMsg)
			break
		}
		var OperandError bool

		for { //поиск арифметического знака и разбиение оставшегося текстого материала по этому знаку ТОЛЬКО на две части: до и после знака
			str2 = "+"
			str1, str3, OperandError = strings.Cut(InputString, str2)
			if OperandError {
				break
			} //если во введенном пользователе массиве текста есть хотя бы один плюс выходим из вложенного for
			str2 = "-"
			str1, str3, OperandError = strings.Cut(InputString, str2)
			if OperandError {
				break
			} //если во введенном пользователе массиве текста есть хотя бы один минус
			str2 = "*"
			str1, str3, OperandError = strings.Cut(InputString, str2)
			if OperandError {
				break
			} //если во введенном пользователе массиве текста есть хотя бы один знак  *

			str2 = "/"
			str1, str3, OperandError = strings.Cut(InputString, str2)
			if OperandError {
				break
			} //если во введенном пользователе массиве текста есть хотя бы один знак /
			if !OperandError {
				break // если не нашлось ни одного валидного арифметического знака выходим для начала из этого for;;
			}
		}

		if !OperandError {
			errorString = ("Учтите, Вы НЕ ввели арифметический знак из ряда { + - * / } между операндами.\nИ скорее всего еще что-то не то ввели\n" + BasicErrorMsg)
			break // это выход уже из глобального for
		}

		str1 = strings.ToUpper(str1)
		str3 = strings.ToUpper(str3)

		if !(checkOperandsArab(str1) || checkOperandsRome(str1)) {
			errorString = ("Ваш как минимум первый операнд не вписывается в рабочий диапазон,\n вводите от 1 до 10 арабскими или римским (от I до X ) цифрами.\n С нулем, отрицательными, 11 и более значениями приложение работать НЕ будет!\n" + BasicErrorMsg)
			break
		}

		if !(checkOperandsArab(str3) || checkOperandsRome(str3)) {
			errorString = ("Ваш второй операнд не вписывается в рабочий диапазон, вводите от 1 до 10 арабскими или римским (от I до X ).\nС нулем, отрицательными, 11 и более значениями приложение работать НЕ будет!\n" + BasicErrorMsg)
			break
		}

		if (checkOperandsRome(str1) && checkOperandsArab(str3)) || (checkOperandsRome(str3) && checkOperandsArab(str1)) {
			errorString = ("Ну Вы даете! У Вас разносортица в операндах. Вводите только римские или только арабские цифры. ") //+ BasicErrorMsg)
			break
		}

		/*rome routines*/
		if checkOperandsRome(str1) {
			op1 = RomeMap[str1]
			op2 = RomeMap[str3]
			switch str2 {
			case "+":
				res = op1 + op2
			case "*":
				res = op1 * op2
			case "/":
				res = op1 / op2
			case "-":
				res = op1 - op2
			}
			if res <= 0 {
				errorString = ("Извините, когда работаете с римскими цифрами, не допускайте результата нулевого, или отрицательного.\nЧтобы получать отрицательные результаты и ноль используйте только арабские цифры") //+ BasicErrorMsg)
				break
			}

			fmt.Println(intToRome100(res)) //rome result

		} else { //arab routines //tests for correct boundaries already done
			op1, _ := strconv.Atoi(str1)
			op2, _ := strconv.Atoi(str3)
			switch str2 {
			case "+":
				res = op1 + op2
			case "*":
				res = op1 * op2
			case "/":
				res = op1 / op2
			case "-":
				res = op1 - op2
			}
			fmt.Println(res) //arab result
		} //arab routines

		str1 = ""
		str2 = ""
		str3 = ""
		InputString = ""
	} //for
	fmt.Println(errorString) //error result prints, then exits program
} //main
