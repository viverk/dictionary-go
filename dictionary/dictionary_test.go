package dictionary

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	// Run tests
	exitCode := m.Run()

	// Exit with the code from tests
	os.Exit(exitCode)
}

func TestDictionary_Add(t *testing.T) {
	d, err := New()
	if err != nil {
		t.Fatalf("Error initializing dictionary: %v", err)
	}
	defer d.Close()

	// Test successful Add
	err = d.Add("testWord", "testDefinition")
	if err != nil {
		t.Errorf("Unexpected error adding word: %v", err)
	}

	// Test Add with invalid data
	err = d.Add("b", "definition")
	if err == nil {
		t.Error("Expected error for short word, but got none")
	}
}

func TestDictionary_Get(t *testing.T) {
	d, err := New()
	if err != nil {
		t.Fatalf("Error initializing dictionary: %v", err)
	}
	defer d.Close()

	// Test Get with an existing word
	d.Add("testWord", "testDefinition")
	result, err := d.Get("testWord")
	if err != nil {
		t.Errorf("Unexpected error getting word: %v", err)
	}
	if result != "testDefinition" {
		t.Errorf("Expected definition 'testDefinition', got '%s'", result)
	}

	// Test Get with a non-existent word
	_, err = d.Get("nonExistentWord")
	if err == nil {
		t.Error("Expected error for non-existent word, but got none")
	}
}

func TestDictionary_Remove(t *testing.T) {
	d, err := New()
	if err != nil {
		t.Fatalf("Error initializing dictionary: %v", err)
	}
	defer d.Close()

	// Add a word to the dictionary
	d.Add("wordToRemove", "definition")

	// Test successful Remove
	err = d.Remove("wordToRemove")
	if err != nil {
		t.Errorf("Unexpected error removing word: %v", err)
	}

	// Test Remove with a non-existent word
	err = d.Remove("nonExistentWord")
	if err == nil {
		t.Error("Expected error for non-existent word, but got none")
	}
}

func TestDictionary_List(t *testing.T) {
	d, err := New()
	if err != nil {
		t.Fatalf("Error initializing dictionary: %v", err)
	}
	defer d.Close()

	// Add some words to the dictionary
	d.Add("word1", "definition1")
	d.Add("word2", "definition2")

	// Test List
	entries, err := d.List()
	if err != nil {
		t.Errorf("Unexpected error listing entries: %v", err)
	}

	// Check if the expected entries are present
	expectedEntries := map[string]string{"word1": "definition1", "word2": "definition2"}
	for word, definition := range expectedEntries {
		if entries[word] != definition {
			t.Errorf("Expected definition '%s' for word '%s', got '%s'", definition, word, entries[word])
		}
	}
}
