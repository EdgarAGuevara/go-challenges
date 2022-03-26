package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
 * Reto #1
 * ¿ES UN ANAGRAMA?
 * Fecha publicación enunciado: 03/01/22
 * Fecha publicación resolución:Ø 10/01/22
 * Dificultad: MEDIA
 *
 * Enunciado: Escribe una función que reciba dos palabras (String) y retorne verdadero o falso (Boolean) según la palabra 1 sea anagrama o no de la segunda palabra.
 * Un Anagrama consiste en formar una palabra reordenando TODAS las letras de otra palabra inicial.
 * NO hace falta comprobar que ambas palabras existan.
 * Dos palabras exactamente iguales no son anagrama.
 *
 * Información adicional:
 * - Usa el canal de nuestro discord (https://mouredev.com/discord) "🔁reto-semanal" para preguntas, dudas o prestar ayuda a la acomunidad.
 * - Puedes hacer un Fork del repo y una Pull Request al repo original para que veamos tu solución aportada.
 * - Revisaré el ejercicio en directo desde Twitch el lunes siguiente al de su publicación.
 * - Subiré una posible solución al ejercicio el lunes siguiente al de su publicación.
 *
 */

func main() {

	var resp bool = isAnagram("I am Lord Valdimort", "Tom Marvolo Riddle")

	if resp {
		fmt.Println("Hello, is Anagram!")
	} else {
		fmt.Println("Hello, is NOT Anagram!")
	}

}

func isAnagram(subject string, result string) bool {

	var trimedSubject = strings.TrimSpace(subject)
	var trimedResult = strings.TrimSpace(result)
	fmt.Printf("Subject: %v  Result: %v \n", trimedSubject, trimedResult)

	var lowerSubject = strings.ToLower(trimedSubject)
	var lowerResult = strings.ToLower(trimedResult)
	fmt.Printf("Lower Subject: %v. lower Result: %v. \n", lowerSubject, lowerResult)

	var chunkedSubject = strings.ReplaceAll(lowerSubject, " ", "")
	var chunkedResult = strings.ReplaceAll(lowerResult, " ", "")
	fmt.Printf("Junked Subject: %v. Junked Result: %v \n", chunkedSubject, chunkedResult)

	var splitedSortedSubject = strings.Split(chunkedSubject, "")
	var splitedSortedResult = strings.Split(chunkedResult, "")
	sort.Strings(splitedSortedSubject)
	sort.Strings(splitedSortedResult)
	fmt.Printf("Splited and Sorted Subject: %v. Splited and Sorted Result: %v \n", splitedSortedSubject, splitedSortedResult)

	var sordtedSubject = strings.Join(splitedSortedSubject, "")
	var sordtedResult = strings.Join(splitedSortedResult, "")
	fmt.Printf("Sorted Subjet: %v  Sorted Result: %v \n", sordtedSubject, sordtedResult)

	return (strings.Compare(sordtedSubject, sordtedResult) == 0)
}
