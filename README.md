# Go Text Search Engine

A fast and efficient full-text search engine implemented in Go, designed to search through Wikipedia titles. The engine uses an inverted index for quick lookups and implements various text processing features like tokenization, stopword removal, and stemming.

## Features

- Fast full-text search using inverted index
- Text processing pipeline:
  - Tokenization
  - Lowercase conversion
  - Stopword removal
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

Specify a custom search query:
```bash
go run main.go -q "Albert Einstein"
```

Use a different data file:
```bash
go run main.go -p "path/to/your/data.gz" -q "search term"
```

### Command-line Options

- `-p`: Path to the gzipped data file (default: "enwiki-latest-all-titles-in-ns0.gz")
- `-q`: Search query (default: "Albert Einstein")

## Project Structure

- `main.go`: Entry point and CLI interface
- `utils/`
  - `documents.go`: Document loading and management
  - `index.go`: Inverted index implementation
  - `tokenizer.go`: Text tokenization
  - `filter.go`: Text processing filters (lowercase, stopwords, stemming)

## Performance

The search engine uses several optimization techniques:
- Inverted index for O(1) term lookups
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