# Go-Stl-Parser

## Project Structure:

```
├── cmd
│   └──main  
|      └──main.go      //flag processing
├── internal
│   └──decode
|      └──decode.go     \\ all the parser logic
|      └──tridimensional.go     \\ data structures and built in functions
├── internal
│   └──stl                \\ stl example files for testing
|      └──cube.stl
|      └──liver.stl
|      └──solid.stl
├── go.mod              \\ Go dependencies file
├── go.sum
├── makefile            
```

## Considerations

### Structure
Simple project structure implemented to parse a stl file and organize the solid information in data structures.

### Design
The parser itself was implemented following a simple _Finite State Machine_ design taking into consideration the relevant possible states while reading the ascii file.

### Potential Improvements
After tackling the parser task, the best approach I could think was to build something that resembles a compiler. Create a step that would generate a stack of tokens and their informations and then process the token using go routines(making the code indisputably more efficient). 

## Running the code
`make run` will run the default case. Which parses the cube.stl file. There is also a fp(filepath) flag which can be used to specify a stl file to be parsed.
In this case, run `make run fp=(path)`
