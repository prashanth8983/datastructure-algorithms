package main

import "fmt"

// ============ STRUCT EMBEDDING (composition over inheritance) ============

// TODO: Define a base struct:
//       type Animal struct {
//           Name string
//       }
//       func (a Animal) Speak() string { return a.Name + " makes a sound" }

// TODO: Define an embedded struct (Dog "inherits" Animal):
//       type Dog struct {
//           Animal          // embedded — Dog gets all Animal methods
//           Breed string
//       }

// TODO: Override Speak on Dog:
//       func (d Dog) Speak() string { return d.Name + " barks" }

// TODO: Define another embedded struct without override:
//       type Cat struct {
//           Animal
//           Indoor bool
//       }
//       // Cat uses Animal's Speak() automatically

// ============ INTERFACE EMBEDDING ============

// TODO: Define small interfaces:
//       type Reader interface { Read() string }
//       type Writer interface { Write(data string) }

// TODO: Embed them into a combined interface:
//       type ReadWriter interface {
//           Reader
//           Writer
//       }

// TODO: Define a struct File that implements ReadWriter:
//       type File struct {
//           Name    string
//           Content string
//       }
//       func (f *File) Read() string          { return f.Content }
//       func (f *File) Write(data string)     { f.Content = data }

func main() {
	// TODO: Create a Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
	// TODO: Print dog.Name (promoted from Animal), dog.Breed, dog.Speak()

	// TODO: Create a Cat{Animal: Animal{Name: "Whiskers"}, Indoor: true}
	// TODO: Print cat.Speak() — uses Animal's Speak since Cat doesn't override

	// TODO: Create a File, Write to it, Read from it
	// TODO: Pass File to a function that accepts ReadWriter interface

	_ = fmt.Println // remove once used
}
