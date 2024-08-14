package media

/*
MEDIA
Calcula la media de un
listado de números.
*/

func Execute(numbers []int) float64 {
	var media float64
	for _, i := range numbers {
		media += float64(i)
	}
	return media / float64(len(numbers))
}
