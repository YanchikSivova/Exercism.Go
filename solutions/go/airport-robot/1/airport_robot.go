package airportrobot
import "fmt"
type Greeter interface{
    LanguageName()string
    Greet(a string) string
}

type Italian struct{}

type Portuguese struct{}

func SayHello(name string, robot Greeter) string{
    return fmt.Sprintf("I can speak %s: %s", robot.LanguageName(), robot.Greet(name))
}

func (it Italian)LanguageName() string{
    return "Italian"
}

func (it Italian)Greet(name string) string{
    return "Ciao " + name + "!"
}

func (port Portuguese) LanguageName() string{
    return "Portuguese"
}

func (port Portuguese) Greet(name string) string{
    return "Olá " + name + "!"
}
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
