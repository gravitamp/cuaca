package main

import (
	"sort"
	"strconv"
)

type sorted struct {
	category    string
	probability float64
}

// Classifier is what we use to classify documents
type Classifier struct {
	cuaca               map[string]map[string]int
	totalWords          int
	categoriesDocuments map[string]int
	totalDocuments      int
	categoriesWords     map[string]int
	threshold           float64
}

// create and initialize the classifier
func createClassifier(categories []string, threshold float64) (c Classifier) {
	c = Classifier{
		cuaca:               make(map[string]map[string]int),
		totalWords:          0,
		categoriesDocuments: make(map[string]int),
		totalDocuments:      0,
		categoriesWords:     make(map[string]int),
		threshold:           threshold,
	}

	for _, category := range categories {
		c.cuaca[category] = make(map[string]int)
		c.categoriesDocuments[category] = 0
		c.categoriesWords[category] = 0
	}
	return
}

// Train the classifier
func (c *Classifier) Train(category string, time string, dmin string, dmax string, tmin string, tmax string) {
	condition := c.switchCond(time, dmin, dmax, tmin, tmax)
	for word, count := range countCuaca(condition) {
		c.cuaca[category][word] += count
		c.categoriesWords[category] += count
		c.totalWords += count
	}
	c.categoriesDocuments[category]++
	c.totalDocuments++
}

//counting
// clean up and split words in document, then stem each word and count the occurrence (tidak pakai stem)
func countCuaca(condition string) (wordCount map[string]int) {
	// cleaned := cleaner.ReplaceAllString(waktu, "")
	// words := strings.Split(cleaned, " ")
	// wordCount = make(map[string]int)
	// for _, word := range words {
	// 	if !stopWords[word] {
	// 		key := stem(strings.ToLower(word))
	// 		wordCount[key]++
	// 	}
	// }
	return
}

//get condition
func (c *Classifier) switchCond(waktu string, dmin string, dmax string, tmin string, tmax string) (category string) {
	// get all the probabilities of each category
	dens_min, _ := strconv.ParseUint(dmin, 10, 32)
	dens_max, _ := strconv.ParseUint(dmax, 10, 32)
	temp_min, _ := strconv.ParseUint(tmin, 10, 32)
	temp_max, _ := strconv.ParseUint(tmax, 10, 32)
	condition := dens_min + dens_max + temp_min + temp_max
	cond := strconv.FormatUint(uint64(condition), 10)

	return cond
}

// Classify a document
func (c *Classifier) Classify(waktu string, dmin string, dmax string, tmin string, tmax string) (category string) {
	// get all the probabilities of each category
	cond := c.switchCond(waktu, dmin, dmax, tmin, tmax)
	prob := c.Probabilities(cond)

	// sort the categories according to probabilities
	var sp []sorted
	for c, p := range prob {
		sp = append(sp, sorted{c, p})
	}
	sort.Slice(sp, func(i, j int) bool {
		return sp[i].probability > sp[j].probability
	})

	// if the highest probability is above threshold select that
	if sp[0].probability/sp[1].probability > c.threshold {
		category = sp[0].category
	} else {
		category = "unknown"
	}

	return
}

// Probabilities of each category
func (c *Classifier) Probabilities(condition string) (p map[string]float64) {
	p = make(map[string]float64)
	for category := range c.cuaca {
		p[category] = c.pCategoryDocument(category, condition)
	}
	return
}

// p (category)
func (c *Classifier) pCategory(category string) float64 {
	return float64(c.categoriesDocuments[category]) / float64(c.totalDocuments)
}

// p (condition | category)
func (c *Classifier) pDocumentCategory(category string, condition string) (p float64) {
	p = 1.0
	for cond := range countCuaca(condition) {
		p = p * c.pWordCategory(category, cond)
	}
	return p
}

func (c *Classifier) pWordCategory(category string, condition string) float64 {
	return 8 //float64(c.cuaca[category][stem(word)]+1) / float64(c.categoriesWords[category])
}

// p (category | condition)
func (c *Classifier) pCategoryDocument(category string, condition string) float64 {
	return c.pDocumentCategory(category, condition) * c.pCategory(category)
}
