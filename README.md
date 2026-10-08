# Notes Tool

A Go learning project for creating, viewing, and managing single-line notes organized into persistent collections.

## Task

Build an interactive command-line tool that opens or creates a notes collection, displays its entries, adds new notes, and deletes notes by number while preserving changes between runs.

## Learning Focus

Practice functions, slices, buffered input, command-line arguments, input validation, and reading and writing plain-text files.

---

## What the Tool Does

**Notes Tool** is an interactive terminal-based manager for quick personal notes, ideas, and tasks. 

Key features include:
- **Collection-Based Organization:** Keeps different topics (e.g., `work`, `shopping`, `ideas`) isolated into independent plain-text collections.
- **Interactive Menu Loop:** Allows performing multiple actions in a single session without restarting the program.
- **Formatted Note Display:** Shows all notes formatted with three-digit numbers (e.g., `001 - note one`).
- **Deletion:** Supports removing notes by line number or typing `0` to safely cancel the operation.
- **Configuration:** Uses simple text files for persistent storage with no external database dependencies.

---

## Architecture & Project Structure

The project logic is modularized across dedicated files:

- `notestool.go`: Main entry point, program greeting, and interactive menu loop.
- `helping_func.go`: CLI argument validation, help message display, and file verification.
- `add.go`: Handles reading note input and appending it to the collection.
- `show.go`: Formats and displays existing notes with three-digit indices.
- `delete.go`: Validates indices and safely removes notes from the collection.
- `storage.go`: Handles disk I/O operations (`LoadNotes` and `SaveNotes`).

---

## Usage

### Compilation & Execution

The application uses the Go standard library without external dependencies. The current `go.mod` declares Go 1.27.1.

Select the files listed below explicitly: `notestool_my_only.go` is an alternative implementation with its own `main` function, so including every Go file causes a duplicate entry point.

Build the binary:
```bash
go build -o notestool notestool.go helping_func.go add.go show.go delete.go storage.go
```

Run the compiled executable:
```bash
./notestool <collection_name>
```

Or run directly using Go:
```bash
go run notestool.go helping_func.go add.go show.go delete.go storage.go <collection_name>
```

---

### Command-Line Arguments & Help Message

The tool takes **exactly one argument**: the name of the collection you want to open or create.

- **Valid invocation:**
  ```bash
  ./notestool coding_ideas
  ```

- **Help message:**  
  If no argument is passed, if more than one argument is provided, or if the argument is `help`, `-h`, or `--help`, the tool displays a usage hint and exits immediately:
  ```text
  $ ./notestool
  Usage: ./notestool <file>
  ```

---

## Interactive Menu Operations

When launched with a valid collection name, the tool welcomes the user and opens an interactive loop with four operations:

1. **`1. Show notes.`**  
   Reads and prints all saved notes in the collection, formatted with leading zeros (`001 - ...`, `002 - ...`).
2. **`2. Add a note.`**  
   Prompts for a line of text, appends it to the collection, and saves it immediately to disk.
3. **`3. Delete a note.`**  
   Prompts for the item number to remove. Entering `0` cancels the action without making changes.
4. **`4. Exit.`**  
   Exits the menu loop and terminates the program cleanly.

After an action is performed, the menu is redisplayed until option `4` is chosen.

---

## Example Walkthrough

```text
$ ./notestool testtag
Welcome to the notes tool!

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
2

Enter the note text:
note three

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two
003 - note three

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
3

Enter the number of note to remove or 0 to cancel:
3

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
4
```

---

## Data Storage

The data storage system is built around plain text files:

- **File Naming:** Each collection corresponds to a separate plain text file bearing the exact name of the collection (e.g., collection `coding_ideas` is stored in a file named `coding_ideas`).
- **Row-Based Records:** Each individual note is stored on a separate row ending with a standard newline character (`\n`).
- **Persistence Across Runs:** 
  - If the collection file does not exist when launched, it is automatically created.
  - If the collection already exists, existing records are loaded into memory at startup.
  - Any additions or deletions rewrite and update the collection file immediately, preserving all changes between program runs.

---

## Technical Constraints & Packages

The project uses these Go standard library packages:
- `bufio`: Buffered reading of terminal input and scanning text files line-by-line.
- `fmt`: Terminal I/O formatting and file printing.
- `os`: File opening, creation, closing, and argument parsing.
- `strconv`: Parsing string inputs into numerical indices.
- `strings`: Trimming whitespace and newlines from user input.

---

## Team & Contribution

- **Nadezda Korepanova (Leader):** `notestool.go`, `helping_func.go` — CLI interface, argument parsing, greeting, and main menu loop.
- **Pavithra Kannan (Member):** `add.go`, `show.go`, `delete.go` — Functions to show, add, and delete notes.
- **Oleksandr Mohutnov (Member):** `storage.go` — Functions to load and save notes to persistent disk storage, boundary checks, and project configuration (`go.mod`, `.gitignore`).
