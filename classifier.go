package main

import (
	"sort"
)

type sorted struct {
	category    string
	probability float64
}

// Classifier is what we use to classify documents
type Classifier struct {
	cuaca               (map[string]map[string]int)
	totalWords          int
	categoriesDocuments map[string]int
	categoriesWords     map[string]int
	threshold           float64
}

// create and initialize the classifier
func createClassifier(categories []string, threshold float64) (c Classifier) {
	c = Classifier{
		cuaca:               make(map[string]map[string]int),
		totalWords:          0,
		categoriesDocuments: make(map[string]int),
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

	c.categoriesWords[category]++
	c.totalWords++
	c.categoriesDocuments[category]++
}

// Classify a document
func (c *Classifier) Classify(waktu string, dmin string, dmax string, tmin string, tmax string) (category string) {
	// get all the probabilities of each category
	prob := c.Probabilities(waktu, dmin, dmax, tmin, tmax)

	// sort the categories according to probabilities
	var sp []sorted //category, prob
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
func (c *Classifier) Probabilities(waktu string, dmin string, dmax string, tmin string, tmax string) (p map[string]float64) {
	p = make(map[string]float64)
	for category := range c.cuaca {
		p[category] = c.pCategoryDocument(category, waktu, dmin, dmax, tmin, tmax)
	}
	return
}

// p (category)
func (c *Classifier) pCategory(category string) float64 {
	return float64(c.categoriesDocuments[category]) / float64(len(train))
}

// p (condition | category)
func (c *Classifier) pDocumentCategory(category string, condition string) float64 {
	return float64(c.cuaca[category][condition]+1) / float64(c.categoriesWords[category])
}

// p (category | condition1|cond2|cond3|cond4)
func (c *Classifier) pCategoryDocument(category string, waktu string, dmin string, dmax string, tmin string, tmax string) float64 {
	return c.pDocumentCategory(category, waktu) * c.pDocumentCategory(category, dmin) *
		c.pDocumentCategory(category, dmax) * c.pDocumentCategory(category, tmin) *
		c.pDocumentCategory(category, tmax) * c.pCategory(category)
}
