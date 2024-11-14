---
subject: ТРПО
title: Лекция GO
author: Беднов А.О.
type: лекция
lang: ru
format:
  revealjs:
    transition: slide
    background-transition: fade
    slide-number: true
    chalkboard:
      buttons: true
    preview-links: auto
    incremental: true
    scrollable: true
---

# Лекция GO

---

## Основы Golang

---

### Структура программы

- Программа на языке Go хранится в одном или нескольких файлах. Каждый файл с программным кодом должен принадлежать какому-нибудь пакету. И вначале каждого файла должно идти объявление пакета, к которому этот файл принадлежит. Пакет объявляется с помощью ключевого слова package.

- В файле может использоваться функционал из других пакетов. В этом случае используемые пакеты надо импортировать с помощью ключевого слова import. Импортируемые пакеты должны идти после объявления пакета для текущего файла:

``` go
package main
import "fmt"
```

---

- Например, в данном случае текущий файл будет находиться в пакете main. И далее он подключает пакет fmt.

- Причем главный пакет программы должен называться "main". Так как именно данный пакет определяет, что будет создаваться исполняемый файл приложения, который после компиляции можно будет запускать на выполнение.

- После подключения других пакетов располагаются объявления типов, переменных, функций, констант.

---

- При этом входной точкой в приложения является функция с именем main. Она обязательно должна быть определена в программе. Все, что выполняется в программе, выполняется именно в функции main.

``` go
package main
import "fmt"

func main() {
    fmt.Println("Hello Go!")
}
```

---

- Базовым элементом программы являются инструкции.
- Например, вызов функции `fmt.Println("Hello Go!")` представляет отдельную инструкцию. Каждая инструкция выполняет определенное действие и размещается на новой строке:

``` go
package main
import "fmt"

func main() {
    fmt.Println("Hello Go!")
    fmt.Println("Hello Golang!")
    fmt.Println("Hello Go!")
}

```

---

- Здесь функция main содержит три инструкции, которые выводит на консоль строку, и каждая из инструкций размещается на новой строке.

- Можно размещать несколько инструкций и на одной строке, но тогда их надо отделять точкой запятой:

``` go
fmt.Println("Hello Go!"); fmt.Println("Hello Golang!"); fmt.Println("Hello Go!")
```

---

- В то же время размещение инструкций на новой строке представляет более читабельный формат, поэтому более предпочтительно для использования.

---

### Комментарии

- Программа может иметь комментарии. Комментарии служат для описания действий, которые производит программа или какие-то ее части. При компиляции комментарии не учитываются и не оказывают никакого влияния на работу приложения. Комментарии бывают однострочными и многострочными.

- Однострочный комментарий располагается в одну строку после двойного слеша ////. Все, что идет после этих символов, воспринимается компилятором как комментарий. Многострочный комментарий заключается между символами /\* и \*/ и может занимать несколько строк:

``` go
/*
    Первая программа
    на языке Go
*/
package main    // определение пакета для текущего файла
import "fmt"    // подключение пакета fmt

// определение функции main
func main() {
    fmt.Println("Hello Go!")    // вывод строки на консоль
}
```

---

### Пакеты

- В Go программы структурированы с помощью пакетов. Основная функция пакета должна быть в . Если вы создаете проект, в котором нет пакета main, то это означает, что вы создаете библиотеку. В сообщении Go говорится, что имя пакета должно совпадать с последним элементом пути импорта. Путь импорта — это, по сути, путь к проекту из рабочего пространства Go. Теперь, когда мы говорим об импорте, go дает нам оператор для импорта определенного пакета либо из стандартной библиотеки, либо из некоторого пакета в нашем рабочем пространстве go.

---

```go
package main

import (
"fmt"
"math"
)

func main() {
fmt.Println("Make the zero value useful.")

fmt.Printf("Square root of 8: %v\n", math.Sqrt(8))
}
```

---

#### Экспорт имен

- Если вы использовали Java, то вы, вероятно, уже получили удовольствие от этих бесконечных определений методов, я говорю о вас. В Go нет _конкретных_ модификаторов доступа, таких как публичный или частный. В Go имена могут быть либо на уровне пакета, либо публичными. И вы делаете это, делая первую букву имени заглавной буквой. При импорте пакета можно ссылаться только на те имена, которые начинаются с заглавной буквы. Например([GoPlay](https://goplay.space/#CrFk7n7AWd6)):`public static void main`

```go
package main

import (
"fmt"
"math"
)

func main() {
// Run this and you will get an error. After that
// change the pi to be Pi and see what happens.
fmt.Println(math.pi)

}
```

---

### Функции

- Функции в Go могут принимать любое количество аргументов и возвращать любое количество результатов. Типичная модель в Go заключается в том, что вы возвращаете желаемые результаты плюс ошибку, но об ошибках мы расскажем позже. Возвращаемые значения функции могут быть названы, если это так, то Go будет рассматривать их как локальные переменные в области видимости функции. Пример ([GoPlay](https://goplay.space/#xPmMJhUXExA)):

```go
package main

import "fmt"

// declare x and y as named return results
func split(sum int) (x, y int) {
x = sum * 4 / 9
y = sum - x
return
}

func main() {
fmt.Println(split(17))
}
```

---

- Кое-что интересное в Go — это вариативные функции. Эти функции могут быть вызваны с любым количеством конечных аргументов. Примером из стандартной библиотеки может быть [fmt. Println](https://golang.org/src/fmt/print.go?s=7595:7644#L253). Вот пример, который проиллюстрирует все это ([GoPlay](https://goplay.space/#9xI5pni7y-g)):

```go
package main

import "fmt"

// Here's a function that will take an arbitrary number
// of `int`s as arguments.
func sum(nums ...int) {
fmt.Print(nums, " ")
total := 0
for _, num := range nums {
total += num
}
fmt.Println(total)
}

func main() {
// Variadic functions can be called in the usual way
// with individual arguments.
sum(1, 2)
sum(1, 2, 3)

// If you already have multiple args in a slice,
// apply them to a variadic function using
// `func(slice...)` like this.
nums := []int{1, 2, 3, 4}
sum(nums...)
}
```

---

### Операторы управления потоком

---

#### For

- Go дает нам только одну конструкцию цикла, цикл. Это означает, что у нас нет времени, или повторения до этого, или чего-то подобного. Циклы for имеют базовую структуру, аналогичную той, что используется в C, за исключением того, что мы не используем скобки, они фактически будут ошибкой компиляции:`for`

```go
for i := 0; i < 2; i++ {
// code
}
```

- Первая и последняя часть for может быть необязательной, что означает, что мы можем использовать только вычислительную часть, и у нас в основном есть время вот так:

```go
for i != 4 {
// do some stuff
}
```

---

- **Range** перебирает элементы в различных структурах данных. Давайте посмотрим, как использовать диапазон с некоторыми структурами данных. ([GoPlay](https://play.golang.org/p/ChWJFN-Zaoy))

```go
package main

import "fmt"

func main() {

    // Here we use `range` to sum the numbers in a slice.
    // Arrays work like this too.
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // `range` on arrays and slices provides both the
    // index and value for each entry. Above we didn't
    // need the index, so we ignored it with the
    // blank identifier `_`. Sometimes we actually want
    // the indexes though.
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // `range` on map iterates over key/value pairs.
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // `range` can also iterate over just the keys of a map.
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // `range` on strings iterates over Unicode code
    // points. The first value is the starting byte index
    // of the `rune` and the second the `rune` itself.
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```

---

#### If

- Если операторы, как for, не используют круглые скобки. Мы можем начать операторы if с оператора, который должен быть выполнен перед условием, например:

```go
if err := funcReturnsError(); err != nil {
// very important stuff
}
```

---

#### Switch

- Оператор `switch` — это более короткий способ записи последовательности операторов if - else. Он запускает первый вариант, значение которого равно выражению условия. Переключатель Go похож на тот, что используется в C, C++, Java, JavaScript и PHP, за исключением того, что Go запускает только выбранный случай, а не все последующие. Это означает, что нам не нужен этот уродливый оператор break в конце каждого случая, как в упомянутых языках. Еще одно важное отличие состоит в том, что случаи switch в Go не обязательно должны быть константами, а соответствующие значения не обязательно должны быть целыми числами. Например([GoPlay](https://goplay.space/#SIsdHJKlgxe))

```go
package main

import (
"fmt"
"runtime"
)

func main() {
fmt.Print("Go runs on ")
switch os := runtime.GOOS; os {
case "darwin":
fmt.Println("macOS.")
case "linux":
fmt.Println("Linux.")
default:
fmt.Printf("%s.", os)
}
}
```

---

#### Defer

- `Defer` — это то, что часто используется в Go. Эти операторы откладывают выполнение функции до тех пор, пока не вернется окружающая функция. Аргументы, которые она получает, оцениваются немедленно, но функция не выполняется до тех пор, пока не вернется окружающая функция. Отложенные вызовы функций помещаются в стек вызовов. Когда функция возвращается, среда выполнения go извлекает каждую из отложенных функций и выполняет их (это структура LIFO). Пример ([GoPlay](https://goplay.space/#6hzIT6lIo5F)):

```go
package main

import "fmt"

func main() {
fmt.Println("counting")

for i := 0; i < 10; i++ {
defer fmt.Println(i)
}

fmt.Println("done")
}
```

---

### Переменные и типы

---

#### Переменные

- Для хранения данных в программе применяются переменные. Переменная представляет именованный участок в памяти, который может хранить некоторое значение. Для определения переменной применяется ключевое слово var, после которого идет имя переменной, а затем указывается ее тип:

``` go
var имя_переменной тип_данных
```

- Имя переменной представляет произвольный идентификатор, который состоит из алфавитных и цифровых символов и символа подчеркивания. При этом первым символом должен быть либо алфавитный символ, либо символ подчеркивания. При этом имена не должны представлять одно из ключевых слов: break, case, chan, const, continue, default, defer, else, fallthrough, for, func, go, goto, if, import, interface, map, package, range, return, select, struct, switch, type, var.

---

- Например, простейшее определение переменной:

``` go
var hello string
```

- Данная переменная называется hello и она представляет тип string, то есть некоторую строку.

- Можно одновременно объявить сразу несколько переменных через запятую:

``` go
var a, b, c string
```

---

- В данном случае определены переменные a, b, c, которые имеют тип string. В этом случае опять же в конце указывается тип данных, и все переменные принадлежат этому типу.

- После определения переменной ей можно присвоить какое-либо значение, которое соответствует ее типу:

``` go
package main
import "fmt"

func main() {
    var hello string
    hello = "Hello world"
    fmt.Println(hello)
}
```

---

- Поскольку переменная hello представляет тип string, ей можно присвоить строку. В данном случае переменная hello хранит строку "Hello world". С помощью функции Println значение этой переменной выводится на консоль.

- Также важно учитывать, что Go - регистрозависимый язык, то есть переменные с именами "hello" и "Hello" будут представлять разные переменные:

``` go
var hello string
hello = "Hello world"
fmt.Println(Hello)  // ! Ошибка переменной Hello нет, есть переменная hello
```

---

- Также можно сразу при объявлении переменной присвоить ей начальное значение. Такой прием называется инициализацией:

``` go
package main
import "fmt"

func main() {
    var hello string = "Hello world"
    fmt.Println(hello)
}
```

---

- Если мы хотим сразу определить несколько переменных и присвоить им начальные значения, то можно обернуть их в скобки:

``` go
package main
import "fmt"

func main() {
    var (
        name string = "Tom"
        age int = 27
    )

    fmt.Println(name)   // Tom
    fmt.Println(age)    // 27
}
```

---

- Отличительной особенностью переменных является то, что их значение можно многократно изменять:

``` go
package main
import "fmt"

func main() {
    var hello string = "Hello world"
    fmt.Println(hello)  // Hello world

    hello = "Hello Go"
    fmt.Println(hello)  // Hello Go

    hello = "Go Go Go Ole Ole Ole"
    fmt.Println(hello)  // Go Go Go Ole Ole Ole
}
```

---

##### Краткое определение переменной

- Также допустимо краткое определение переменной в формате:

``` go
имя_переменной := значение
name := "Tom"
```

---

- После имени переменной ставится двоеточие и равно и затем указывается ее значение.

``` go
package main
import "fmt"

func main() {
    name := "Tom"
    fmt.Println(name)

}
```

- В этом случае тип данных явным образом не указывается, он выводится автоматически из присваиваемого значения.

---

#### Типы

- Go имеет почти все типичные значения типов:

``` go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

---

- Этими типами можно манипулировать с помощью различных операторов. Строки могут быть добавлены вместе с целыми числами и числами с плавающей точкой.
- Логические операторы имеют все логические операторы, как и ожидалось.`+`

---

- Эти типы используются при объявлении переменных. В Go вы можете объявлять переменные тремя различными способами, которые будут показаны в коде ниже. Go также позволяет объявлять константы, ставя ключевое слово перед именем. Выводится тип константы. [Запустите его онлайн](https://goplay.space/#cmQu-3Uf58J)`const`

```go
package main

func main() {
// Declare with no initialization, go will give the default value
// which in this case is 0.
var someInt int
// Giving it a specific initial value
var someBool, hi = true, "hi"
// Inferring the type from the right side of the expression
hello := "hello world" // hello will be of type string

// declaring constants. Type gets inferred
const number = 2;
const str = "some string";

fmt.Printf("someInt: %v\n", someInt)
fmt.Printf("someBool: %v\nhi: %v\n", someBool, hi)
fmt.Printf("helloWorld: %v\n", hello)
fmt.Printf("number: %v\n", number)
fmt.Printf("str: %v\n", str)
}
```

---

- Когда вы объявляете переменную в Go без инициализации, эта переменная будет содержать нулевое значение указанного типа. Например, в случае , его нулевое значение будет равно .
- Если вы хотите узнать больше о том, почему Go использует этот формат для объявления переменных и параметров, вы можете ознакомиться с [синтаксисом объявления Go](https://blog.golang.org/gos-declaration-syntax)`someInt``0``bool``false``""``<name> <type>`

---

- Приведение типов в Go выполняется путем заключения переменной или значения, к которому вы хотите привести, в круглые скобки, перед которыми указан тип, к которому вы хотите привести его. Например([GoPlay](https://goplay.space/#rAUMn61amkI)):

```go
package main

func main() {
someInt := 2
castedFloat := float64(someInt)
fmt.Printf("type of someint: %T\ntype of castedFloat: %T\n", someInt, castedFloat)
// will result in:
// type of someint: int
// type of castedFloat: float64
}
```

---

### Упражнение

Прежде чем переходить к остальному содержимому, перейдите к [циклам и функциям, потренируйтесь](https://tour.golang.org/flowcontrol/8) и реализуйте то, что там требуется. Попробуйте запустить код локально, чтобы убедиться, что ранее настроенная среда работает правильно.

---

### Указатели

У Го есть указатели. Указатель в Go будет содержать адрес значения в памяти. Нулевое значение указателей это . Для объявления указателя на тип мы используем синтаксис, аналогичный C:`nil`

---

В данном случае это указатель на int. Чтобы получить доступ к значению, на которое указывает указатель, или к адресу памяти заданной переменной, когда можно сделать это ([GoPlay](https://goplay.space/#5_eTXtJ1h43)):`somePointer`

```go
package main

func main() {
var someInt *int
otherInt := 2
pointer := &otherInt
fmt.Printf("zero-value of someInt: %v\n", someInt)
fmt.Printf("value pointer points to: %v\n", *pointer)
}
```

---

### Структуры

Чтобы сгруппировать поля вместе, Go определяет , эти структуры очень похожи на те, которые определены в языке C. Они определяются с помощью , каждое поле будет иметь свой собственный тип, и они могут быть инициированы различными способами. Например([GoPlay](https://goplay.space/#x4njww69awl)):`structs``type name struct`

```go
package main

import "fmt"

type someStruct struct {
intField  int
boolField bool
}

func main() {
// this will have the zero-value for all its fields
s := someStruct{}
// fields can get initialized by order
s = someStruct{2, true}
// we can name the fields we are initializing
l := someStruct{
boolField: false,
intField:  2,
}

// We can then access the fields of the struct using the
// dot notation:
fmt.Printf("l bool field: %v\n", l.boolField)
fmt.Printf("s int field field: %v\n", s.intField)
}
```

---

Если вы объявляете указатель на структуру, то можно подумать, что для доступа к полям нам придется использовать обозначение указателя. Но Go - это хорошо, и нам не нужно вводить эти 3 дополнительных символа, мы можем просто сделать следующее ([GoPlay](https://goplay.space/#wKurxTznpp5)):

**ПРИМЕЧАНИЕ:** обратите внимание, что в приведенном выше примере атрибуты структуры являются приватными, потому что первая буква написана в нижнем регистре.

---

```go
package main

type someStruct struct {
intField int
boolField bool
}

func main() {
p := &someStruct{2, false}
fmt.Printf("doing p.intField works: %v\n", p.intField)
fmt.Printf("but also we can do (*p).boolField: %v\n", (*p).boolField)
}
```

---

- **Анонимные структуры** — это то, что дает нам Go, что очень удобно при тестировании. В принципе, мы можем сгруппировать поля без указания для них типа. Например([GoPlay](https://goplay.space/#H8CneBlRDfh)):

```go
package main

func main() {
// see that it has no name that can be used to reference
// the "type" of this structure
s := struct {
x int
y int
}{
x: 2,
y: 3,
   }
fmt.Printf("x: %v y: %v", s.x, s.y)
}
```

---

### Композиция над наследованием

Как мы уже говорили ранее во введении. Го – это все про композицию, на самом деле никакого наследования не существует. Если вы хотите поделиться поведением и данными чего-то, вы компонуете что-то с этим типом, нет необходимости в подклассах и указании сложного дерева иерархии классов и наследования. Придя из мира ООП, поначалу это может показаться странным. Но композиция на самом деле является хорошо понятной концепцией ООП, и в Го вы будете использовать ее довольно часто. Чтобы узнать немного больше о том, как это работает, вы можете прочитать [этот раздел](https://golang.org/doc/effective_go.html#embedding) Effective Go, чтобы увидеть несколько реальных примеров того, как вы можете извлечь из этого пользу. Давайте посмотрим на простой пример того, как вы можете использовать это ([GoPlay](https://goplay.space/#AnaEldEzcBl)):

```go
package main

import "fmt"

// User holds information of a given user.
type User struct {
ID             int
Name, Location string
}

// Here you can see that player embeds the User type.
// This way we are saying that a Player holds is composed
// of a User and a GameID.
type Player struct {
User
GameID int
}

func main() {
p := Player{}
p.ID = 42
p.Name = "Globant"
p.Location = "La Plata"
p.GameID = 90404
fmt.Printf("%+v", p)
// This will print the following:
// {User:{ID:42 Name:Globant Location:La Plata} GameID:90404}
// You can see that the User type is embedded in the Player
// structure.
}
```

---

### Массивы, срезы и карты

- В Go есть кое-что, что часто сбивает с толку новичков, и это разница между массивами и срезами.
- **Массивы** — это, как вы могли бы подумать, список индексированных значений фиксированного размера, которые имеют один и тот же тип.
- **Ломтики** – это нечто немного более интересное. Первое отличие заключается в том, что массивы имеют фиксированный размер, в то время как срезы выделяются динамически. На практике вы, вероятно, всегда будете видеть срезы вместо массивов.
- Стоит отметить, что сами срезы не содержат массив данных.
- Они просто содержат три поля: Len, Cap и Data.
- Len будет длиной массива,
- Cap — максимальной емкостью массива,
- а Data — указателем на массив поддержки.
- [Здесь](https://golang.org/pkg/reflect/#SliceHeader) вы можете увидеть структуру среза. Срезы и массивы инициализируются по-разному и позволяют нам выполнять различные типы операций. Ниже приведен код, который объяснит все это более подробно ([GoPlay](https://goplay.space/#NPh97D1qgEY)):`structs`

```go
package main

import "fmt"

func main() {
// This is how you would declare an array
// var someArray [3]int
// Slices can be declared in many different ways.
// This slice will have its zero value that is going to be nil, since
// no array has been alocated yet.
// var names []string
// slice of strings with an initial size of 2 and unlimited capacity
// otherNames := make([]string, 2)
// slice of strings with an initial size of 2 and maximum capacity of 4
// capacity := make([]string, 2, 4)
// We can initialize with values
numbers := []int{1, 2, 3, 4, 5}
// Slices let us do operations using the indices
oneToThree := numbers[0:2]
fmt.Println(oneToThree)
// We can omit one of the indices and it will go to the last or the first
threeToFive := numbers[2:]
fmt.Println(threeToFive)
fmt.Println(numbers)

// Incrementing the size of slice.
// If we want to append values to already declared slice we can use the append function
oneToThree = append(oneToThree, 4)
fmt.Println(oneToThree)
// append doesn't care about the receiver, so we could declare a new variable with append:
fiveToSix := append(threeToFive[len(threeToFive)-1:], 6)
fmt.Println(fiveToSix)
}
```

---

Чтобы узнать больше об использовании срезов и о том, как они работают внутри компании, обратитесь к [этой записи блога](https://blog.golang.org/go-slices-usage-and-internals).
**Упражнения:** Перейдите в [Go Tour Slices потренируйтесь](https://tour.golang.org/moretypes/18) и реализуйте то, что там требуется.

---

Карты в основном похожи на словари в Python или HashMaps в Java. Они сопоставляют ключ с заданным значением и позволяют нам получать доступ к этим значениям с помощью указанных ключей. Нулевое значение карты будет равно нулю, как и в случае со срезами. Оба ключа и значения могут быть любого заданного типа, от структур до базовых типов и интерфейсов (о которых мы расскажем позже). Так же, как и срезы, мы можем использовать для создания карт. Примеры ([GoPlay](https://goplay.space/#xjeZ48dknCc)):`make`

```go
package main

import "fmt"

type someStruct struct {
intField  int
boolField bool
}

func main() {
// this a nil map since it's the zero-value
// var m map[string]someStruct
// declaring and initializing maps
m := map[string]someStruct{
"key1": someStruct{2, false},
"key2": someStruct{3, true},
}
// We can reference the fields of the struct by accessing the value of the map
fmt.Printf("m[\"key1\"].intField=%v\n", m["key1"].intField)
// We can allocate new key-value pairs
m["key3"] = someStruct{4, true}
fmt.Printf("m[\"key3\"].boolField=%v\n", m["key3"].boolField)
}
```

---

**Упражнение:** перейдите к [упражнению Go tour maps](https://tour.golang.org/moretypes/23) и выполните то, что требуется.

---

### Типы функций

В Го функции являются гражданами первого класса. Это означает, что вы можете использовать функции так, как если бы они были такими же, как и любые другие типы (ints, string, что у вас есть). На самом деле, наряду со стандартной библиотекой и многими другими пакетами вы найдете очень много типов функций. Например, в библиотеке есть тип функции, который обычно используется под названием [http. HandlerFunc](https://godoc.org/net/http#HandlerFunc) с сигнатурой . Чтобы объявить и использовать типы функций, вы можете сделать что-то вроде этого ([GoPlay](https://goplay.space/#rZ4Z48yEykG)):`net/http``func (http.ResponseWriter, *http.Request)`

```go
package main

import (
"errors"
"fmt"
)

func main() {
var fun func() error
fun = func() error {
fmt.Println("hi from fun")
return errors.New("error from fun")
}
someFunc(fun)
}

func someFunc(f func() error) error {
fmt.Println("hi from someFunc")
// return the result of calling f
return f()
}
```

---

**Упражнение:** перейдите к [упражнению по закрытию тура Go](https://tour.golang.org/moretypes/26) и выполните то, что требуется.

---

### Методы

- Go, в отличие от объектно-ориентированного программирования, не имеет классов.
- Но у него есть методы, которые вы можете определить для конкретных типов.
- Это может показаться немного странным, но метод — это всего лишь функция со специальным _приемником_.
- Этот приемник может быть **любого** типа, целыми числами, структурами, строками и всем, что вы можете придумать (они также могут быть определены для типов функций, крутое начало, правда?).
- Поскольку ресивер может быть любого типа, он может иметь определенное состояние. Например, если у меня есть структура с двумя полями int, то эти два поля будут доступны из методов, которые я определил поверх этой структуры.
- Если вы хотите изменить эти состояния и хотите, чтобы они отражались за пределами метода, вам придется определить получатель как получатель указателя, это означает, что получатель будет содержать адрес состояния, которым вы манипулируете.
- Если вместо этого мы будем использовать приемники значений, то методы будут работать с копиями исходных значений.
- Для удобства Go позволит нам вызывать методы, у которых есть указатели-получатели со значением, так что нам не придется вводить 3 дополнительных символа (doing вместо ).
- Прекратите говорить и покажите мне код ([GoPlay](https://goplay.space/#sHBFKyoY5wf)):`v.Method()``(&v).Method()`

```go
package main

import "fmt"

type mstr string

func (m *mstr) Hi() {
// print the value of the receiver
fmt.Printf("Hi %v\n", *m)
}

func main() {
// we can call hi using a value, go will translate to
// a memory address since that is what Hi expects as the
// receiver
name := mstr("Gopher")
name.Hi()

// we can call hi too when we have an actual memory address
otherName := &name
otherName.Hi()
}
```

---

### Интерфейсы

- Интерфейсы позволяют нам группировать функции с определенными указанными нами сигнатурами.
- Значением интерфейсов может быть любое значение, реализующее эти методы.
- В Go интерфейсы реализованы неявно.
- Этого нет или чего-то подобного.
- Если данный тип реализует функции, определенные данным интерфейсом, то этот тип будет реализовывать этот интерфейс без нашего участия. Интерфейсы содержат значение и тип, значение которого является значением конкретного конкретного конкретного типа, о котором идет речь.
- Одна интересная вещь в значениях интерфейса в Go заключается в том, что они могут быть nil, так что по сути вы можете вызвать метод для значения nil, который реализует интерфейс, и это будет нормально, в других языках это, вероятно, приведет к исключениям нулевого указателя. Вы можете задаться вопросом: что произойдет, если в интерфейсе нет методов?
- Разве ни один конкретный тип не реализовал бы этот интерфейс?? ... Что ж, вы правы.
- Пустой интерфейс, на который обычно ссылаются as, может содержать значения любого типа. Иногда это полезно при работе со значениями неизвестного типа. Например, в пакете стандартной библиотеки у вас есть метод [Encoder.Encode](https://godoc.org/encoding/json#Encoder.Encode), который получает данные, которые вы хотите закодировать, через него, так что вы можете отправить в этот метод все, что захотите.
- **БУДЬТЕ ОСТОРОЖНЫ** при использовании пустого интерфейса, не злоупотребляйте им, всякий раз, когда вы определяете API для своей библиотеки, попытайтесь увидеть, можете ли вы вместо этого использовать пользовательский интерфейс с методами, которые выражают значение типов, которые будет обрабатывать API. При определении интерфейсов Go рекомендует уменьшить количество методов, которые вы помещаете в этот интерфейс, так как **чем больше интерфейс, тем слабее абстракция**. Вот и было много текста, пришло время для некоторого кода ([GoPlay](https://goplay.space/#FNh0cTZqrTb)):`implements``interface{}``json``interface{}`

```go
package main

import "fmt"

// Note the format. Typically one method interfaces have
// a the name of the interface be the name of the method
// + er. For example, the interface that has the method
// Write is called Writer, in this case the method name
// is greet so the interface will be called Greeter.
type Greeter interface {
Greet(name string)
}

type str struct{}

// str type implements the Greeter interface implicitly.
func (str) Greet(name string) {
fmt.Printf("Hi %v\n", name)
}

func main() {
// s is of type str so it implements Greeter
s := str{}
s.Greet("gopher")

// whatType receives an empty interface so we can
// send whatever we want to the function
whatType(s)
whatType(3)
whatType("hi")

// greetAndBye expects a Greeter, so whatever type
// that implements the Greet(name string) method
// can be sent to this function
greetAndBye(s, "gopher")
}

func whatType(v interface{}) {
fmt.Printf("type of %v: %T\n", v, v)
}

func greetAndBye(g Greeter, name string) {
g.Greet(name)
fmt.Printf("Bye %v\n", name)
}
```

---

Вы собираетесь реализовать решение двух различных упражнений, которые покажут вам интерфейсы из стандартной библиотеки, которые обычно используются. **Упражнение:** перейдите к [упражнению стрингеров Go tour](https://tour.golang.org/methods/18) и выполните то, что требуется. **Упражнение:** перейдите к [упражнению читателей Go tour](https://tour.golang.org/methods/22) и выполните то, что требуется.

---

#### Утверждения типов

- Есть вещь, которую вы можете сделать с интерфейсами, которая называется утверждением типов. Это обеспечивает доступ к базовому конкретному значению интерфейса. По сути, это означает, что при наличии интерфейса вы можете извлечь значение определенного типа. Например([GoPlay](https://goplay.space/#jOqaZ9T9TlU)):

```go
package main

import "fmt"

type Greeter interface {
Greet(name string)
}

type str struct{}

func (str) Greet(name string) {
fmt.Printf("Hi %v\n", name)
}

func main() {
// s is of type str so it implements Greeter
s := str{}
s.Greet("gopher")
extract(s)
}

func extract(v interface{}) {
// if v holds a value of type str then greeter will
// have that value. If it does not then this will triger
// a panic
greeter := v.(str)
greeter.Greet("gopher1")

// instead of triggering a panic we can use a second
// variable that will hold a boolean specifying if v
// is of type str
gr, ok := v.(str)
if !ok {
fmt.Println("v is not of type str")
return
}
gr.Greet("gopher2")
}
```

---

- Возможно, вы задаетесь вопросом, а что, если мы хотим выполнять утверждение типов с помощью нескольких типов, а не только одного? Что ж, поехали. Вы можете использовать переключатели типа ([GoPlay](https://goplay.space/#pomXwCJ9XzA)):

```go
package main

import "fmt"

type Greeter interface {
Greet(name string)
}

type str struct{}

func (str) Greet(name string) {
fmt.Printf("Hi %v\n", name)
}

func main() {
// s is of type str so it implements Greeter
s := str{}
extract(s)
extract(2)
extract("hello")
extract(2.4)
}

func extract(v interface{}) {
switch t := v.(type) {
case str:
t.Greet("GOPHER")
case int:
fmt.Printf("sent an int: %v\n", t)
case string:
fmt.Printf("sent a string: %v\n", t)
default:
fmt.Printf("unsupported type: %T\n", t)
}
}
```

---

### Ошибки

- Ошибки являются предметом большого обсуждения в Golang, в основном потому, что он сильно отличается от того, что мы обычно используем в таких языках, как Java или Ruby.
- Если вы хотите выразить ошибку на этих языках, вы должны создать свое собственное исключение и дать ему осмысленное имя и, возможно, описание.
- В этом случае вам понадобятся дополнительные структуры управления для обработки этих ошибок. В Go мы рассматриваем ошибки как простые значения. Для этого у нас есть тип, который представляет собой не что иное, как [интерфейс с одним методом](https://godoc.org/builtin#error):`try..catch` `error`

```go
type error interface {
Error() string
}
```

---

- В go функции обычно возвращают ошибки, поэтому, если что-то пошло не так, вы просто проверяете, равна ли ошибка нулю или нет. Это означает, что всякий раз, когда вы создаете свои собственные функции, вы должны возвращать все, что хотите, и ошибка, если ничего не пошло не так, просто возвращает nil, в противном случае возвращает ошибку, выражающую проблему. Вы увидите много кода, который выглядит следующим образом:

```go
// Handling errors
err := callFunc()
if err != nil {
// handle the error in here
}

// Returning errors.
// Return the data you want + an error
func importantFunc() (string, error) {
s, err := importantOperation()
if err != nil {
return "", err
}
return s, nil
}
```

**Упражнение:** перейдите к [упражнению Go tour](https://tour.golang.org/methods/20) error exercise и выполните то, что требуется.

---

### Параллелизм

---

#### Горутины

- Если вы слышали о Go, то вы, вероятно, также слышали, что у Go есть потрясающая встроенная поддержка параллелизма.
- Он предельно прост и очень эффективен.
- Для этого в Go используется концепция _горутин,_ горутина — это облегченный поток, управляемый средой выполнения Go.
- Это означает, что в одном потоке ОС у нас может быть несколько горутин, тысячи, если хотите. Эти программы выполняются в одном и том же адресном пространстве, поэтому вам придется быть осторожным при доступе к данным из нескольких программ. Это немного показывает, как вы можете использовать горутины ([GoPlay](https://goplay.space/#cRTTaMlqWj0)):

```go
package main

import (
"fmt"
"time"
)

func importantFeature() {
fmt.Println("doing important work")
}

func main() {
importantFeature()

// you can spawn new goroutines with a simple
// go funcName
go importantFeature() // this now created a brand new goroutine

s := "inside goroutine"
// you can also spawn goroutines using an
// anonymous functions
go func(someString string) {
fmt.Println("printing from the goroutine")
}(s)

for i := 0; i < 5; i++ {
// we can spawn as many as we like. Try changing
// i<5 to something like i<1000
go func(i int) {
time.Sleep(5 * time.Millisecond)
fmt.Printf("i=%v\n", i)
}(i)
}
// wait a bit so that we don't exit immediately
// we need to wait so that at least some of the
// goroutines get executed. Go doesn't automatically
// wait for all goroutines, if we don't do synchronization
// then we will exit and goroutines won't finish executing.
time.Sleep(25 * time.Millisecond)
}
```

---

#### Каналами

- Для связи и синхронизации между различными горутинами в Го была введена концепция _каналов_. Каналы — это каналы, которые соединяют параллельные горутины. Вы можете отправлять значения из одной горутины и получать их из другой. Например([GoPlay](https://goplay.space/#vteRsUOcf1w)):

```go
package main

import (
"fmt"
"time"
)

func main() {
messages := make(chan string)

go func() {
time.Sleep(1 * time.Second)
// The information flows in the direction of the arrow
// In this case we are sending data into the messages
// channel
messages <- "hi from the goroutine"
}()

// now we are receiving the information that is sent to
// the channel.
// Note that this will block for a second until the
// channel has something we can receive.
m := <-messages
fmt.Println(m)
}
```

---

```go
package main

import (
"fmt"
"time"
)

func main() {
messages := make(chan string, 2)

// we send two messages that will make
// the channel full.
messages <- "hi"
messages <- "bye"

// spawn a new goroutine that will read
// from the channel after three seconds
go func() {
time.Sleep(3 * time.Second)
fmt.Println(<-messages)
}()

// if we try to add a message to the channel
// before one or all of the messages are consumed
// then the operation will block since there
// is no more space for that message to fit.
// So after the previous go routine reads, we
// will be able to send "hi again"
messages <- "hi again"

fmt.Println(<-messages)
fmt.Println(<-messages)
}
```

---

- Мы можем сделать что-то подобное для синхронизации выполнения между горутинами. Например, мы можем использовать блокирующий receive для ожидания, пока горутина не завершит выполнение ([GoPlay](https://goplay.space/#-BGwwrmU1ve)):

```go
package main

import "fmt"
import "time"

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
fmt.Print("working...")
time.Sleep(time.Second)
fmt.Println("done")

// Send a value to notify that we're done.
done <- true
}

func main() {

// Start a worker goroutine, giving it the channel to
// notify on.
done := make(chan bool, 1)
go worker(done)

// Block until we receive a notification from the
// worker on the channel.
<-done
}
```

---

- В предыдущем примере мы показали, что вы можете отправлять каналы в качестве параметров.
- Что вы также можете сделать, так это указать _направление_ канала, это ограничит операции, которые мы можем выполнять. Таким образом, по сути, мы можем сказать: «Это канал только для отправки», «Это канал только для приема» или «Вы можете делать все, что хотите, быть свободным». Давайте перепишем предыдущий пример, обратите внимание, что внутри функции worker мы только отправляем данные в канал, поэтому вместо этого мы можем сделать следующее:

```go
func worker(done chan<- bool) {
fmt.Print("working...")
time.Sleep(time.Second)
fmt.Println("done")

// Send a value to notify that we're done.
done <- true
}
```

---

Если бы мы хотели отправить в функцию канал только для приема, мы могли бы сделать следующее:

```go
// note the direction of the arrow
func someFunc(done <-chan bool) {
// do work
}
```

---

- Закрытие канала указывает на то, что значения по нему больше не будут отправляться. Иногда это может быть полезно для того, чтобы указать на завершение тем, кто получает от него.
- Чтение из закрытого канала не возвращает ошибку, оно возвращает нулевое значение типа канала. Таким образом, если вы объявили канал строк и читаете из него, то после его закрытия вы получите все пустые строки. Если вы используете структуры, то нулевое значение будет равно nil. Это позволяет нам использовать структуру для перебора значений канала до тех пор, пока канал не будет закрыт. Следующий пример будет охватывать следующие варианты использования ([GoPlay](https://goplay.space/#WGVtmD3Ck5C)):`for range`

```go
package main

import (
"fmt"
"time"
)

func main() {
messages := make(chan string, 2)
done := make(chan bool, 1)

go func() {
for {
// this double assignment allows us to check
// if the channel was closed or not
m, ok := <-messages
if !ok {
// if ok is false it means that we cannot
// receive from the channel so it was closed
fmt.Println("channel closed!")
// close the channel so that it sends the nil
// value which will make the last sentence
// of this program ends and the program exits
close(done)
return
}
fmt.Printf("received: %v\n", m)
}
}()
// wait before closing the channel
time.Sleep(2 * time.Second)
messages <- "hi"
messages <- "bye"
close(messages)
<-done
}
```

- Go предоставляет нам выписку под названием .
- Этот оператор позволяет горутине ожидать нескольких операций связи и выполнять действие всякий раз, когда мы получаем что-либо из любого из указанных каналов.
- Пример ([GoPlay](https://goplay.space/#xP74KERfHXl)):

```go
package main

import "time"
import "fmt"

func main() {
c1 := make(chan string)
c2 := make(chan string)

// Each channel will receive a value after some amount
// of time, to simulate e.g. blocking RPC operations
// executing in concurrent goroutines.
go func() {
time.Sleep(1 * time.Second)
c1 <- "one"
}()
go func() {
time.Sleep(2 * time.Second)
c2 <- "two"
}()

// We'll use `select` to await both of these values
// simultaneously, printing each one as it arrives.
for i := 0; i < 2; i++ {
select {
case msg1 := <-c1:
fmt.Println("received", msg1)
case msg2 := <-c2:
fmt.Println("received", msg2)
}
}
}
```

**Упражнение:** перейдите к [упражнению двоичного дерева Go tour](https://tour.golang.org/concurrency/7) и реализуйте то, что требуется.

---

#### Мьютексы

Вы видели, что каналы отлично подходят для общения между разными горутинами, мы даже можем использовать их для синхронизации. Но что, если мы хотим защитить переменную, чтобы только одна горутина могла получить к ней доступ одновременно. Было бы довольно громоздко реализовывать решение, используя только каналы. Что ж, go снова хорош и предоставляет нам библиотеку под названием , которая содержит набор утилит, полезных для взаимного исключения с использованием обычно называемого . Этот пример был взят из GoTour ([GoPlay](https://goplay.space/#7L8oJihNE1G)):`sync``mutex`

```go
package main

import (
"fmt"
"sync"
"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
v   map[string]int
mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
c.mux.Lock()
// Lock so only one goroutine at a time can access the map c.v.
c.v[key]++
c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
c.mux.Lock()
// Lock so only one goroutine at a time can access the map c.v.
defer c.mux.Unlock()
return c.v[key]
}

func main() {
c := SafeCounter{v: make(map[string]int)}
for i := 0; i < 1000; i++ {
go c.Inc("somekey")
}

time.Sleep(time.Second)
fmt.Println(c.Value("somekey"))
}
```

---

**Упражнение:** перейдите к [упражнению веб-краулера Go tour](https://tour.golang.org/concurrency/10) и реализуйте то, что требуется.

---

### Рекапитуляция

Вы просто читаете кучу текста, поэтому давайте зададим набор вопросов, что-то вроде подведения итогов, чтобы проверить ваши знания. В этом списке есть не только вопросы, но и упражнения, которые вам нужно решить. Мы рекомендуем вам поместить все это в репозиторий, чтобы инструкторы могли видеть, как идут дела, сделать небольшой обзор кода и помочь вам с любыми трудностями, с которыми вы можете столкнуться.

- Как бы вы написали заявление в Go?`while`
- Что делает ключевое слово?`defer`
- Поддерживает ли Go указатели? Как передаются аргументы (по значению или по ссылке)?
- Массивы в Go имеют фиксированную длину? А как насчет ломтиков?
- Допустим, у вас есть карта: , как бы вы выполнили поиск и проверили, содержит ли карта значение ключа, который вы искали?`map[string]int`
- Как Go структурирует программы? В чем разница между библиотекой и выполняемой программой?
- Как сделать функцию или тип общедоступными? И как сделать его приватным?
- Вы будете строить простой калькулятор с 4 основными операциями (сложение, вычитание, умножение и деление). Сначала создайте библиотеку, которая предоставляет эти 4 метода. После этого реализуйте программу, которая считывает из командной строки операцию и выводит результат (вызвав ранее реализованную библиотеку). Операцию следует читать так, как вам нравится, но для простоты ограничьтесь 2 операндами и 1 операционным символом. Что-то вроде .`./program 1 + 2`
- Предположим, вы создаете веб-сервер, которому нужна база данных, способная выполнять набор простых операций. Вы знаете, что требования к тому, какую БД использовать, будут меняться. Кроме того, теперь для целей тестирования будет проще не настраивать что-то вроде MySQL. Как бы вы решили эту проблему, используя функцию, которую предоставляет Go?
- Как бы вы построили простую функцию, которая _может принимать любой_ тип аргумента и выводит если этот аргумент имеет примитивный тип. Ограничьтесь только , , и .`int``string``float``bool`
- Как определяются ошибки в Go?
- Хорошо, теперь вы знаете, как определяются ошибки в Go. Время для сборки простого пакета, позволяющего собирать ошибки, указывающие на тип ошибки, ограничьтесь 3 видами: , и . Затем предоставьте в этом пакете функцию, чтобы пользователи могли проверить, является ли ошибка именно той ошибкой, которую они беспокоятся. **ПРИМЕЧАНИЕ: помните,** что не стоит нарушать способ определения ошибок в Go, воспользуйтесь этим.`errors``Internal``ThirdParty``Other`
- Что вы используете, чтобы сделать две функции параллельными?
- Как бы вы синхронизировали две параллельные функции?
- Напишите программу с тремя функциями. Вы будете отправлять что-то (что угодно) по каналу каждую секунду, а другой будет получать и распечатывать это. Третья функция сообщит двум другим функциям остановиться и вернуться (это может быть основной функцией) через 5 секунд. **ПРИМЕЧАНИЕ:** программа не может быть завершена до тех пор, пока отправитель и получатель не вернутся.
