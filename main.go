package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// 	struct Element1			{struct Element2			struct Element3
// 	{ char Addr[50];			char Name[50];			{ char Name[50];
// 		char Name[50];			char Addr[50];…				char Addr[50];
// 		double Price;			};							double Price;
// 		};													};

//ввод из текстового потока (файла) массива данных (operator >>)
//вывод в текстовый поток (файл) и на консоль массива данных (operator <<)
// вывод в бинарный поток (файл) массива данных или один\несколько с позиции k элемент\ов  массива (ф. seekp)
// ввод из бинарного потока (файла) массив данных или один\несколько с позиции k элемент\ов массива (ф. seekg), только после того как записали данные в бинарный файл
// перегрузка оператора ввода и вывода для Struct (operator >>, operator <<)
// operator=
// доступ к элементу (operator[])
// создать па основе этих данных  новый массив (функц. 1 согласно индивидуальному варианту).
// Упорядочить полученный массив в порядке (функц. 3  согласно индивидуальному варианту).

// 2) Для структур (классов) Element1,  Element2,  Element3:
// конструкторы, деструктор
// перегрузка оператора ввода и вывода для Struct (operator >>, operator <<) (можно как ф. шаблон)
// operator=, operator<, operator== (можно  и как функция- шаблон)

// 3) Х.function1(K,M); //Функция-шаблон  внутри класса-шаблона - объединения(или пересечения, или разности)  массивов (согласно инд. варианту)

// 4) Создать отдельную функцию шаблон, с параметром класс-шаблон MASSIV и элемент структуры  :
// search_function2 ( Х  , st) (функц. 2 согласно индивидуальному варианту)

// 5) Реализовать вне класса - шаблона одну дружественную функцию и любой один метод класса шаблона.

// 6) В шаблоне MASSIV, обрабатывать только массив, поля структур не использовать!

// 7) В main должны быть объекты шаблона MASSIV для типов :double (или char или float или long), для трех типов структур (классов): Element1,  Element2,  Element3 и вызваны методы поиска и сортировки и слияния массивов (согласно инд. варианту)
func main() {

	m1 := (&MASSIV[Element1]{}).New()
	m1.TakeFromTextfile("example.txt")

	fmt.Println(m1.data_about_students)
}

type Element1 struct {
	Addr  string
	Name  string
	Price int
}
type Element2 struct {
	Addr string
	Name string
}
type Parser[T Element1 | Element2] interface {
	ToStruct(s string) *T
}

func (e *Element1) ToStruct(s string) *Element1 {
	var current string
	output := make([]string, 0)

	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " || i+1 == len(s) {
			output = append(output, current)
			current = ""
			continue
		}
		current += string(s[i])
	}

	value, err := strconv.Atoi(output[2])
	if err != nil {
		log.Fatal(err)
	}

	return &Element1{
		Addr:  output[0],
		Name:  output[1],
		Price: value,
	}
}

func (e *Element2) ToStruct(s string) *Element2 {
	parts := strings.Split(s, " ")

	return &Element2{
		Addr: parts[0],
		Name: parts[1],
	}

}

type MASSIV[T Element1 | Element2] struct {
	data_about_students []T //каждый элемент - структура 1 или 2
}

func (m *MASSIV[T]) New() *MASSIV[T] {
	return &MASSIV[T]{
		data_about_students: make([]T, 0),
	}
}
func (m *MASSIV[T]) AppendToMASSIV(s *T) {
	m.data_about_students = append(m.data_about_students, *s)
}

// консоль + в файл   (распарсить структуру)
func (m *MASSIV[T]) TakeFromTextfile(textfile string) { //взять из текста массив структур и захерачит в структуру
	var stroka string

	var zero T
	p, ok := any(&zero).(Parser[T]) //any(x).(SomeInterface) — проверяет, реализует ли x интерфейс SomeInterface
	if !ok {
		log.Fatal(fmt.Printf("тип %T не реализует Parser", zero))
	}

	file, err := os.Open(textfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mass := make([]byte, 15)
	for {
		//i := 0
		_, err := file.Read(mass)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(mass); i++ {
			if mass[i] != '\n' {
				stroka += string(mass[i]) //есть строка, нужно распарсить
			} else {
				elem := p.ToStruct(stroka)
				m.AppendToMASSIV(elem) //сюда передать из ToStruct
				stroka = ""
				continue
			}
		}
	}
}

// из файла
func (m *MASSIV[T]) DeliverToTextfile() {

}

// func (m *MASSIV[T]) WriteText(w io.Writer) error {
// 	for _, item := range m.data_ab_students {
// 		if _, err := fmt.Fprintf(w, "%v\n", item); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

//вывод в бинарный поток (файл) массива данных или один\несколько
// с позиции k элемент\ов  массива (ф. seekp)

// func First[T any](s []T) T {
// 	return s[0]
// }
