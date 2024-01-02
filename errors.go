package main

import "fmt"

func catch(err error, throw int, note string) {
	if err != nil {
		if note != "" {

			if throw == 0 {
				note += " " + err.Error()
			}

			fmt.Println(fmt.Errorf(note + "\n"))
		}

		if throw == 1 {
			panic(err)
		}
	}
}
