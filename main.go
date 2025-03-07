package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/adiboy-23/go-text-search-engine/utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-all-titles-in-ns0.gz", "wiki titles file path")
	flag.StringVar(&query, "q", "Albert Einstein", "search query")
	flag.Parse()
	log.Println("Full text search is in progress...")
	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("loaded %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Title)
	}

}
