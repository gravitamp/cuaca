package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// parameters
var (
	testPercentage = 0.1 //presentasi data test
	datafile       = "data-cuaca.csv"
	threshold      = 1.1
	//exampleif `threshold` is `1.5` this means the category with the highest probability
	// needs to be 1.5 times higher than the second highest probability.
	// If the top category fails the threshold we will classify it as `unknown`.
)

var categories = []string{"Hujan", "Berawan", "Cerah"}

// datasets
type document struct {
	time  string
	class string
	dmin  string
	dmax  string
	tmin  string
	tmax  string
}

//dipisahkan untuk training dan test
var train []document
var test []document

func main() {
	setupData(datafile)
	fmt.Println("Data file used:", datafile)
	fmt.Println("no of docs in TRAIN dataset:", len(train))
	fmt.Println("no of docs in TEST dataset:", len(test))

	//buat classifier dengan parameter yang ada
	c := createClassifier(categories, threshold)

	// train on train dataset
	for _, doc := range train {
		c.Train(doc.class, doc.time, doc.dmin, doc.dmax, doc.tmin, doc.tmax)
	}
	// validate on test dataset
	count, accurates, unknowns := 0, 0, 0
	for _, doc := range test {
		count++
		sentiment := c.Classify(doc.time, doc.dmin, doc.dmax, doc.tmin, doc.tmax)
		if sentiment == doc.class {
			accurates++
		}
		if sentiment == "unknown" {
			unknowns++
		}
	}
	fmt.Printf("Accuracy on TEST dataset is %2.1f%% with %2.1f%% unknowns",
		float64(accurates)*100/float64(count), float64(unknowns)*100/float64(count))
	// validate on the first 100 docs in the train dataset
	count, accurates, unknowns = 0, 0, 0
	for _, doc := range train[0:100] {
		count++
		sentiment := c.Classify(doc.time, doc.dmin, doc.dmax, doc.tmin, doc.tmax)
		if sentiment == doc.class {
			accurates++
		}
		if sentiment == "unknown" {
			unknowns++
		}
	}
	fmt.Printf("\nAccuracy on TRAIN dataset is %2.1f%% with %2.1f%% unknowns",
		float64(accurates)*100/float64(count), float64(unknowns)*100/float64(count))

}

func setupData(file string) {
	rand.Seed(time.Now().UTC().UnixNano())
	data, err := readLines(file)
	if err != nil {
		fmt.Println("Cannot read file", err)
		os.Exit(1)
	}
	for _, line := range data {
		s := strings.Split(line, ",")
		waktu, class, dens_min, dens_max, temp_min, temp_max := s[0], s[1], s[2], s[3], s[4], s[5]

		//dibagi data train dan test
		if rand.Float64() > testPercentage {
			train = append(train, document{class, waktu, dens_min, dens_max, temp_min, temp_max})
		} else {
			test = append(test, document{class, waktu, dens_min, dens_max, temp_min, temp_max})
		}
	}
}

// read the file line by line
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
