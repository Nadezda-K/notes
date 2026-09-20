package main

import (
	"path/filepath"
	"reflect"
	"testing"
)

// TestAddNote verifies adding notes to empty and non-empty slices.
func TestAddNote(t *testing.T) {
	var notes []string

	// 1. Add first note
	notes = AddNote(notes, "Buy milk")
	if len(notes) != 1 || notes[0] != "Buy milk" {
		t.Fatalf("AddNote failed on empty slice: got %v", notes)
	}

	// 2. Add second note
	notes = AddNote(notes, "Call mom")
	expected := []string{"Buy milk", "Call mom"}
	if !reflect.DeepEqual(notes, expected) {
		t.Errorf("AddNote failed on non-empty slice: got %v, want %v", notes, expected)
	}
}

// TestDeleteNote validates 1-based index removal and edge cases.
func TestDeleteNote(t *testing.T) {
	initial := []string{"First", "Second", "Third"}

	tests := []struct {
		name     string
		num      int
		expected []string
	}{
		{
			name:     "Delete first item (num = 1)",
			num:      1,
			expected: []string{"Second", "Third"},
		},
		{
			name:     "Delete middle item (num = 2)",
			num:      2,
			expected: []string{"First", "Third"},
		},
		{
			name:     "Delete last item (num = 3)",
			num:      3,
			expected: []string{"First", "Second"},
		},
		{
			name:     "Boundary: zero (num = 0)",
			num:      0,
			expected: []string{"First", "Second", "Third"},
		},
		{
			name:     "Boundary: negative index (num = -1)",
			num:      -1,
			expected: []string{"First", "Second", "Third"},
		},
		{
			name:     "Boundary: index greater than length (num = 99)",
			num:      99,
			expected: []string{"First", "Second", "Third"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Copy slice to prevent mutation of the initial array
			notesCopy := append([]string(nil), initial...)
			got := DeleteNote(notesCopy, tc.num)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("DeleteNote(%v, %d) = %v; want %v", initial, tc.num, got, tc.expected)
			}
		})
	}
}

// TestStorage_SaveAndLoad validates file writing and reading.
func TestStorage_SaveAndLoad(t *testing.T) {
	// Create an isolated temporary directory cleaned up automatically
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test_notes.txt")

	notesToSave := []string{
		"Note 1: Read documentation",
		"Note 2: Write unit tests",
		"Note 3: Commit changes",
	}

	// 1. Test SaveNotes
	SaveNotes(tempFile, notesToSave)

	// 2. Test LoadNotes on the saved file
	loadedNotes := LoadNotes(tempFile)
	if !reflect.DeepEqual(loadedNotes, notesToSave) {
		t.Fatalf("LoadNotes returned %v; want %v", loadedNotes, notesToSave)
	}

	// 3. Test LoadNotes on a non-existent file
	missingFile := filepath.Join(tempDir, "missing.txt")
	emptyNotes := LoadNotes(missingFile)
	if len(emptyNotes) != 0 {
		t.Errorf("LoadNotes(missingFile) = %v; want empty slice", emptyNotes)
	}
}

// TestCheckArguments validates valid command line argument parsing.
func TestCheckArguments(t *testing.T) {
	args := []string{"./notestool", "my_collection.txt"}
	got := CheckArguments(args)
	want := "my_collection.txt"
	if got != want {
		t.Errorf("CheckArguments(%v) = %q; want %q", args, got, want)
	}
}
