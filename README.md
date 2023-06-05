# API clinet to chucknorris.io
This client provides the possibility to get a random joke about Chuck Norris from the resource https://api.chucknorris.io/

## Installation

~~~~
go get github.com/viking311/chucknorris_api_client 
~~~~

## Usage

To create new client instance you need to call the method NewClient().
~~~~
NewClient(timeout time.Duration)
~~~~

The client provides only method GetRandomeJoke wich returns a random joke from https://api.chucknorris.io/
~~~~
GetRandomeJoke() (*Joke, error)
~~~~

The method loggin its request into stdout.

## Example of usage
~~~~
	client, err := chucknorris.NewClient(5 * time.Second)
	if err != err {
		log.Fatal(err)
	}
	joke, err := client.GetRandomeJoke()
	if err != err {
		log.Fatal(err)
	}

	fmt.Println(joke)
~~~~
