package main

import (
    "fmt"
        "regexp"
	)

	func get_words_from(text string) []string{
	    words:= regexp.MustCompile("\\w+")
	        return words.FindAllString(text, -1)
		}

		func count_words (words []string) map[string]int{
		    word_counts := make(map[string]int)
		        for _, word :=range words{
			        word_counts[word]++
				    }
				        return word_counts;
					}

					func console_out (word_counts map[string]int){
					    for word, word_count :=range word_counts{
					            fmt.Printf("%v %v\n",word, word_count)
						        }
							}

							func main() {
							    text := "I am learning Go! Go is a nice language to learn."
							        console_out(count_words(get_words_from(text)))
								}
