package main

// Dative returns the dative case for a given Hungarian word.
func Dative(word string) string {
  for _, r := range word {
    switch r {
      case 'a', 'á', 'o', 'ó', 'u', 'ú':
      return word + "nak"
    }
  }
  return word + "nek"
}