


# Learn Golang

This repository provides practical Go code examples for beginners and intermediates, covering a wide range of language features and concepts.

## Topics Covered

- **Basics**: Getting started with Go, printing output, basic syntax (`1_basics/`)
- **Simple Values**: Data types like integers, strings, booleans, floats (`2_simple_value/`)
- **Variables**: Declaration, type inference, reassignment (`3_variables/`)
- **Constants**: Defining and using constants (`4_constant/`)
- **For Loop**: Looping constructs and patterns (`5_for_loop/`)
- **If-Else**: Conditional statements (`6_if_else/`)
- **Switch**: Switch-case statements (`7_switch/`)
- **Arrays**: Fixed-size collections (`8_array/`)
- **Slices**: Dynamic arrays and operations (`9_slice/`)
- **Maps**: Key-value data structures (`10_maps/`)
- **Range**: Iterating over collections (`11_range/`)
- **Functions**: Function declaration, parameters, return values (`12_functions/`)
- **Variadic Functions**: Functions with variable number of arguments (`13_varadic_function/`)
- **Closures**: Anonymous functions and capturing variables (`14_Closures/`)
- **Pointers**: Memory addresses and dereferencing (`15_Pointer/`)
- **Structs**: Custom data types and fields (`16_struct/`)
- **Struct Embedding**: Composition and inheritance-like behavior (`17_struct_Embeding/`)
- **Interfaces**: Defining behavior and polymorphism (`18_Interface/`)
- **Enums**: Using iota for enumerated constants (`19_enums/`)
- **Generics**: Type-parameterized functions and data structures (`20_generics/`)
- **Goroutines**: Concurrency with lightweight threads (`21_Goroutines/`)
- **Channels**: Communication between goroutines (`22_Channels/`)
- **Mutex**: Synchronization primitives (`23_Mutex/`)
- **File Handling**: Reading and writing files (`24_FileHandling/`)
- **REST API Example**: A complete RESTful API project in Go (`RestAPI/`)

## REST API Project

The `RestAPI/` folder contains a sample RESTful API built with Go. It demonstrates:
- Project structure with `cmd/`, `internal/`, and `config/` folders
- Configuration management (`config/local.yaml`)
- Database usage (`storage/storage.db`)
- API endpoints for managing student data (`cmd/student_api/main.go`)

To run the REST API:
```bash
cd RestAPI
go run cmd/student_api/main.go -config config/local.yaml
```

## Usage

To run any topic example, navigate to its folder and use:
```bash
go run main.go
```

## Requirements

- Go 1.18+

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve examples or add new topics.

## License

MIT License
