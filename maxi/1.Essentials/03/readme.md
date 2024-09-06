## if statement

1.  The basic syntax of an **if** statement in Go is:

```go
if condition {
    // code to execute if condition is true
}
```

notice: The condition must be a boolean expression. Parentheses around the condition are optional and typically omitted in Go.

2.  Go allows you to include a short statement before the condition. This is often used to set up a variable for use in the condition:

```go
if x := someFunction(); x < 10 {
    // use x
}
```

notice: The variable **x** is scoped to the **if** block and any associated **else** blocks.

3. You can use **else** and **else if** for additional conditions:

```go
if condition1 {
    // code for condition1
} else if condition2 {
    // code for condition2
} else {
    // code if no conditions are true
}
```

4. Go uses standard comparison operators:

-   == (equal to)
-   != (not equal to)
-   < (less than)
-   \> (greater than)
-   <= (less than or equal to)
-   \>= (greater than or equal to)

Go uses the following logical operators:

-   && (AND)
-   || (OR)
-   ! (NOT)

```go
if x > 0 && x < 10 {
    // x is between 0 and 10
}

if !isValid {
    // isValid is false
}

```

6. Go uses short-circuit evaluation for logical operators. In **a && b**, if **a** is false, **b** is not evaluated. In **a || b**, if **a** is true, **b** is not evaluated.

7. Variables declared in an **if** statement's initialization are scoped to the **if** block and any associated **else** blocks:

```go
if x := 10; x > 5 {
    fmt.Println(x) // x is accessible here
}
// x is not accessible here

```

8. Go has a common idiom for error handling using if:

This pattern checks if a function returned an error and handles it immediately.

```go
if err := someFunction(); err != nil {
    // handle error
    return err
}
// continue with normal execution

```

notice:in Go, the { must be on the same line as the if statement. This is enforced by the compiler:

```go
// Correct
if x > 0 {
    // code
}

// Incorrect, will not compile
if x > 0
{
    // code
}

```

## for (loop)

1.  In Go, for is the only looping construct. It can be used in several ways to achieve different types of loops.

2.  Three-Component Loop: This is similar to C-style for loops:

```go
for initialization; condition; post {
    // loop body
}
// example
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

```

3. While-like Loop: You can use for as a while loop by only specifying a condition:

```go
for condition {
    // loop body
}
//example
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

4. Infinite Loop: You can create an infinite loop by omitting the condition:

```go
// infinity
for {
    //some code
}
// but you can use break to stop a loop btw we have continue too break take out of a loop and continue skip a step
for {
    // loop body
    if someCondition {
        break
    }
    if otherCondition {
        continue
    }
}
```

5. Nested Loops: You can nest loops within each other:

```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Printf("(%d, %d)\n", i, j)
    }
}

```

### Additional Notes:

-   Unlike other languages, there's no do-while loop in Go.
-   The post statement in a for loop (e.g., **i++**) is executed at the end of every iteration, including one that ends with a **break** statement.
-   You can omit parts of the for loop syntax. For example, **for ; condition ;** is valid.

```go
// While-like loop
    j := 0
    for j < 3 {
        fmt.Printf("While-like: %d\n", j)
        j++
    }
```

## swaitch

1. The basic syntax of a switch statement in Go looks like this:

```go
switch expression {
case value1:
    // code block
case value2:
    // code block
default:
    // code block
}

func main() {
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("It's Monday")
    case "Tuesday":
        fmt.Println("It's Tuesday")
    default:
        fmt.Println("It's another day")
    }
}

func main() {
    day := "Saturday"
    switch day {
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        fmt.Println("It's a weekday")
    case "Saturday", "Sunday":
        fmt.Println("It's the weekend")
    }
}


```

2. The switch expression is evaluated once, and its value is compared against the values in each case.

3. You can have multiple cases for a single switch statement.

4. The default case is optional and is executed when no other case matches.

5. Unlike in some other languages, Go automatically breaks after each case. You don't need to explicitly use the **break** keyword.

6. If you want execution to continue to the next case, you can use the **fallthrough** keyword.

```go
func main() {
    num := 5
    switch num {
    case 5:
        fmt.Println("Five")
        fallthrough
    case 6:
        fmt.Println("Six")
    case 7:
        fmt.Println("Seven")
    }
}

```

7. You can specify multiple values in a single case, separated by commas.

8. Go supports a special form of switch called a type switch, which allows you to switch on the type of an interface value.

```go
func printType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %v\n", v)
    case string:
        fmt.Printf("String: %v\n", v)
    default:
        fmt.Printf("Unknown type\n")
    }
}

```

9. You can omit the expression in a switch statement, which is equivalent to **switch true**.

```go
func main() {
    hour := 14
    switch {
    case hour < 12:
        fmt.Println("Good morning")
    case hour < 17:
        fmt.Println("Good afternoon")
    default:
        fmt.Println("Good evening")
    }
}

```

10. Cases can include expressions, not just constant values

## writeFile

you can write a file in go with **os** package. in **os** package we have a method call **WriteFile** for writing a file. this **WriteFile** accept 3 parameter. first is **name** the name is your file name and the path of the file **fileName.txt** or **./test/fileName.txt** (in this setuation go cant make directory so you should do it). second parameter is the data but this data should be a collection of byte(data []byte) so we can give string but befor that we should use fmt.Sprint() to make them for callection of byte. the 3th parameter is permission noting special but i write all code possible:

```go
func writeBalanceToFile(balance float64) {
	var balanceText = fmt.Sprint(balance)
	os.WriteFile("./test/balance.txt", []byte(balanceText), 0644)
}
```

### permission code

Code Octal Description \
0000 0o000 No permissions \
0100 0o100 Owner execute \
0200 0o200 Owner write \
0300 0o300 Owner write and execute \
0400 0o400 Owner read \
0500 0o500 Owner read and execute \
0600 0o600 Owner read and write \
0700 0o700 Owner read, write, and execute \
0001 0o001 Others execute \
0002 0o002 Others write \
0003 0o003 Others write and execute \
0004 0o004 Others read \
0005 0o005 Others read and execute \
0006 0o006 Others read and write \
0007 0o007 Others read, write, and execute \
0010 0o010 Group execute \
0020 0o020 Group write \
0030 0o030 Group write and execute \
0040 0o040 Group read \
0050 0o050 Group read and execute \
0060 0o060 Group read and write \
0070 0o070 Group read, write, and execute \
0111 0o111 Owner, group, and others execute \
0222 0o222 Owner, group, and others write \
0333 0o333 Owner, group, and others write, execute \
0444 0o444 Owner, group, and others read \
0555 0o555 Owner, group, and others read, execute \
0666 0o666 Owner, group, and others read, write \
0777 0o777 Owner, group, and others read, write, execute

## readFile

with **os.ReadFile** you can read the file this mathod just take the path(name) of the file and return data and error(if we have a error) so you can have your data:

```go
func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data) // your data is a collection of byte []byte so you need to convet that to a string
	balance, _ := strconv.ParseFloat(balanceText, 64) // in this application we work on we need float so we use strconv string conversion to parse our string to a float
	return balance
}
```

notice : ParseFloat converts the string s to a floating-point number with the precision specified by bitSize: 32 for float32, or 64 for float64. When bitSize=32, the result still has type float64, but it will be convertible to float32 without changing its value.

## error handling

```go
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("can not read file, so you can go with 1000 bounus!")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("can not convert string to float, so you can go with 1000 bounus!")
	}
	return balance, nil
}

var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		fmt.Println("_____________")
	}
```

the code below we findout how to handel that func return error with the value and we use if to handle that and we figureout how to make our function return error with our data and how to handle that

## panic

```go
if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		fmt.Println("_____________")
		panic("Can't continue, sorry.")
	}

```

```
output:
    ERROR:
    can not read file, so you can go with 1000 bounus!
    _____________
    panic: Cant continue, sorry.

    goroutine 1 [running]:
    main.main()
        D:/go/go-training/maxi/Essentials/03/bank.go:83 +0x785
    exit status 2
```

Panic is used to abort the normal flow of a program when an unrecoverable error occurs. It's typically used for programmer errors or unexpected situations that shouldn't occur in normal operation.

When to Use Panic:

-   Unrecoverable errors (e.g., failing to initialize a crucial component)
-   Programming errors (e.g., accessing an array out of bounds)
-   Situations where the program can't continue safely

Panic Behavior: When a panic occurs:

-   The current function execution stops
-   Any deferred functions are run
-   The program crashes with a stack trace if not recovered
