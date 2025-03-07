# Go Text Search Engine

A fast and efficient text search engine implemented in Go, designed to search through Wikipedia article titles. The engine uses an inverted index for quick lookups and implements various text processing features like tokenization, stopword removal, and stemming.

## Features

- Fast title search using inverted index
- Text processing pipeline:
  - Tokenization (splits text on non-letter and non-number characters)
  - Lowercase conversion
  - Stopword removal (common words like "a", "the", "in", etc.)
  - Word stemming (using Snowball stemmer)
- Support for compressed data files (gzip)
- Command-line interface

## Prerequisites

- Go 1.16 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/adiboy-23/go-text-search-engine.git
cd go-text-search-engine
```

2. Install dependencies:
```bash
go mod tidy
```

3. Download the Wikipedia titles dataset:
```bash
wget https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-all-titles-in-ns0.gz
```

## Usage

Run the search engine with default parameters:
```bash
go run main.go
```

Search for specific terms:
```bash
go run main.go -q "Albert Einstein"
```

### Command-line Options

- `-p`: Path to the gzipped titles file (default: "enwiki-latest-all-titles-in-ns0.gz")
- `-q`: Search query (default: "Albert Einstein")

## Project Structure

- `main.go`: Entry point and CLI interface
- `utils/`
  - `documents.go`: Document loading and title management
  - `index.go`: Inverted index implementation and search functionality
  - `tokenizer.go`: Text tokenization and analysis
  - `filter.go`: Text processing filters (lowercase, stopwords, stemming)

## How It Works

1. **Document Loading**: Reads compressed Wikipedia titles and assigns each title a unique ID
2. **Indexing**: 
   - Processes each title through the text analysis pipeline
   - Creates an inverted index mapping tokens to document IDs
3. **Searching**:
   - Processes the search query through the same text analysis pipeline
   - Finds documents containing all search terms using intersection algorithm
   - Returns matching document IDs and titles

## Performance Features

- Inverted index for O(1) token lookups
- Efficient intersection algorithm for multi-term queries
- Memory-efficient document storage
- Streaming processing of gzipped files

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Snowball Stemmer](https://github.com/kljensen/snowball) for word stemming
- Wikipedia for providing the titles dataset 