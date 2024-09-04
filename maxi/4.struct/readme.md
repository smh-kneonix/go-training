## struct

```go
type user struct {
	firstName string
	lastName  string
	birtyear  string
	createdAt time.Time
}
```

struct is a collection of data for better maintain in our code. look at **createdAt** we can add another struct to our struct

```go
    userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")


	var userData user
	userData = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birtyear:  userBirthdate,
        createdAt: time.Now(),
	}
    fmt.Println(userData)
    {fardin kamali 11/27/1999 {13956118768162509588 14520949901 0x980ec0}}
```

NOTICE: if you put your struct litral(userData is a struct litral,that create from a struct) in order that struct created:
first: firstName string
second: lastName string
3th: birtyear string
4th: createdAt time.Time
when you wanna create a struct litral no need to pass key:

```go
var userData user
	userData = user{
		userFirstName,
		userLastName,
		userBirthdate,
        time.Now(),
	}
```

if you change order userFirstName,userLastName in the code below you dident get error but your code is wrong

notice2: if you dident pass one or more value of a struct that should become a nullish value of that type you specific like if userFirstName is string that be "" or for number that become 0

```go
var userData user
	userData = user{
		userFirstName,
		userBirthdate,
	}

output:
fardin  11/27/1999 {0 0 <nil>}}
```

so now we can use pointer to avoid create a copy in our memore look at this example:

```go
var userData user
	userData = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birtyear:  userBirthdate,
		createdAt: time.Now(),
	}

	outputUserDetail(&userData)

func outputUserDetail(user *user) {
	fmt.Printf(`username is %s and the userlastName is %s,this user born in %s. this info created at %s`, user.firstName, user.lastName, user.birtyear, user.createdAt.Local())
}
```

notice: mabey you wandring about why when we use **user** in function as a pointer we dident use **\*** befor **user**, actully this is a shortcut that go provide, you know for useing struct as a pointer you should act like **(\*user).firstName** but this is too much so we use shortcut

## method to struct

we can attach function to a struct, the function that attach to a struct we call them method:

```go
func (u user) outputUserDetail(arg1 string) {
	fmt.Printf(`%s: username is %s and the userlastName is %s,this user born in %s. this info created at %s`, arg1 ,u.firstName, u.lastName, u.birtyear, u.createdAt.Local())

    userData.outputUserDetail("test")
}
```

for this you open prantesis in front on func keyword befor function name and pass the name or struct user in this case, if you wanna access to this struct you can create a varible in this case **u** for accessing to that struct items you can pass your own argument to this method its up to you if you want

### mutate struct litral with method

```go
func (u *user) clearUser() {
	u.firstName = ""
	u.lastName = ""
}
```

notice: for mutate struct litral you should pass pointer (address) of your struct litral not just copy that we do it in section method to struct (we can do it like pointer, that can be kinda complex data and we can avoid save copy in memory) but for mutate we need point to struct litral items to change them

### create new user like constructor function

```go
func newUser(firstName, lastName, birthdate string) user {
	return user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}
var userData user
	userData = newUser(userFirstName, userLastName, userBirthdate)
```

you can avoid copy:

```go
func newUser(firstName, lastName, birthdate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}
var userData *user
	userData = newUser(userFirstName, userLastName, userBirthdate)
```

## exporting from other package

```go
// in user package
type User struct {
	FirstName string
	LastName  string
	birthdate string
	createdAt time.Time
}

// in main package
var userData *user.User
	userData = &user.User{
		FirstName: userFirstName,
	}
```

we meet this before but even you should make items exportable too :) .
but its not right that items accessble from outside of user package so we make our constractore exportable and use that for create struct

notice: its not commen but you see in most pattern like errors, that for constructor we choose New for name of method so keep in mind

## embeded struct

in other language you have class and inharitance but go we dont have so for this we use embeding struct in other struct

```go
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// embeded struct
type Admin struct {
	email string
	password string
	// you can pass a name like
	// user User
	// but you can do it anonimusly like this
	User
}
```

notice :keep in mind if you pass a custom name that all with lowercase you can not access to other method and items in that User struct and you shuld always use this name like admin.user.firstname

```go
type Admin struct {
	email string
	password string
	user User
}

// in main
admin := Admin.NewAdmin("test","passtest")
admin.user.outputUserDetail()
```

## custom type

we can create our own custom type like alise or ... for example

```go
type str string

func main(){
    var name str
}
```

but last example like so useless lets create some usefull

```go
type str string

func (text str) log(){
    fmt.Println(text)
}
func main(){
    var name str = "fardin"
    name.log() // fardin
}
```

## bufio

actully fmt.Scan method is just for a single word or character and you get mess up when trying to have some space or line break or ... so for this perpose we use bufio that another package that help us to get user input from terminal

```go
func GetInput(prompt string) string {
	fmt.Print(prompt)

    // this line create a reader for reading something and os.stdin tell reader you should be acept a standard input from os package(terminal)
	reader := bufio.NewReader(os.Stdin)

    // this line tell reader you should accept a string for ending you reading work and we pass \n(enter) you should use qoute not double
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("invalid input")
		return ""
	}
    // with strings.TrimSuffix we remove \n or in windows some time for linebreake \r  for a fresh we wanna text
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

```

## json

you can convert your struct to a json with **encoding/json** package and **Marshal** method and you can add tag to your struct items for cosumize json property this is built-in feature like **`json:"title"`**.

notice: if your items in struct not first letter capital cant create a json and your json is like { } :(

```go
import "encoding/json"

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

```
