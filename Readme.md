# CRAPLang - CSS Really is a Programming Language
> not to be confused with crablang, which is a completely different thing.

CRAP is a general purpose programming language inspired by CSS.

## Getting Started
- No Comments, who needs them anyway. just write self-documenting code.
- Selectors are used to define functions.
- Attributes are used to define parameters.
- Custom properties are used to define variables.
- declarations are function calls with values as parameters.
- `@if`, `@elif`, `@else`, `@return` are used for control flow.
- Nested selectors are supported.
- `print` is used to print values to the console.
- `()` is used as a placeholder for default values in function calls.
If there is no default value, the parameter is required.
- No null values but also no optional types either, because...sir this is CRAP.
- Check out the examples below for more information.

## Usage
- Clone the repository
```bash
git clone https://github.com/shreyassanthu77/crap.git
cd crap
```
- Make sure you have go installed
- Build main.go
```bash
go build -o crap main.go
```

- Run the examples
```bash
./crap examples/*.css
```

## Docker
- Build the docker image
```bash
docker build -t crap . && docker run -it crap
```
this will open a shell inside the container, you can run the examples from there.
- Run the examples
```bash
./crap examples/*
```

## Examples

### Hello World
```css
main {
    print: "Hello World!";
}
```

### Factorial
```css
factorial[n] {
    @if $n == 0 || $n == 1 {
        @return 1;
    }

    @return $n * factorial($n - 1);
}

main {
    print: factorial(5);
}
```

### Fibonacci
```css
fibonacci.rec[n][a=0][b=1] {
	@if $n == 0 {
		@return $a;
	}
	print: $a;
	@return fibonacci.rec($n - 1, $b, $a + $b);
}

fibonacci[n] {
	@return fibonacci.rec($n, (), ());
}

main {
	fibonacci: 10;
}
```
## Progress

- [x] Lexer
- [x] Parser
- [x] Interpreter
- [ ] Compiler
