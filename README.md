
---

# Ascii-reverse

**Ascii-reverse** is a Go-based command-line application designed to reverse the process of converting the graphic representation into a text. Given a graphic representation of a string in ASCII art, the program will convert it back into its original text form.

## Objectives

This project is a continuation of an earlier ASCII art project (ASCII-Art-Output), but with a twist: instead of generating ASCII art, the program reverses the process by converting ASCII art back into text. 

The program accepts a text file containing the ASCII art of a string and converts it back to the original string.

### Features

- **Flag-Based Argument**: The program accepts a flag `--reverse=<fileName>`, where `<fileName>` is the name of the file containing the ASCII art. It then prints the original string in normal text.
- **Strict Flag Format**: The `--reverse=<fileName>` flag must be used exactly as shown. Any deviation from this format will trigger a usage message.

## Instructions

- **Programming Language**: The project is written in Go.
- **Testing**: The project includes unit tests to verify the functionality of individual components.

## Installation

1. Clone the repository:
   ```
   $ git clone https://learn.zone01kisumu.ke/git/danyonyi/ascii-art-reverse.git
   ```

2. Navigate into the project's directory:
   ```
   $ cd ascii-art-reverse
   ```

## Usage

To use the `Ascii-reverse` program, follow these steps:

1. **Create an ASCII Art File**:
   - Create a text file containing ASCII art that represents a string. For example, `file.txt` could contain the following:

    ```plaintext
     _              _   _          
    | |            | | | |         
    | |__     ___  | | | |   ___   
    |  _ \   / _ \ | | | |  / _ \  
    | | | | |  __/ | | | | | (_) | 
    |_| |_|  \___| |_| |_|  \___/  
                                  
                                  
    ```

2. **Run the Program**:
   - Execute the program with the `--reverse` flag followed by the file name:

    ```bash
    $ go run . --reverse=file.txt
    ```

   - The program will output:

    ```plaintext
    hello
    ```

3. **Usage Instructions**:
   - If the flag format is incorrect or missing, the program will display the usage message:

    ```plaintext
    Usage: go run . [OPTION]

    Example: go run . --reverse=<fileName>
    ```

## Code Overview

Here’s a brief explanation of the main functions:

- **`main()`**: Entry point of the program. It handles argument parsing and triggers the reversal process.
- **`ReadFile(fileName string) ([]string, bool)`**: Reads the content of the provided file and splits it into lines.
- **`SliceFile(fileSlice []string) [][]string`**: Splits the file content into slices representing each character in ASCII art.
- **`CheckPattern(char, word []string) bool`**: Compares a segment of the ASCII art with a known pattern to identify the corresponding character.
- **`TrimFound(length int, word []string) []string`**: Trims the identified character from the ASCII art to continue processing the next character.

## Conclusion

The `Ascii-reverse` project is a reverse-engineering exercise that converts ASCII art back into text. This project demonstrates proficiency in Go, as well as the ability to handle string manipulation, file reading, and pattern matching.

---
