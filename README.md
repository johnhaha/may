# may

## Install

```bash
go get github.com/johnhaha/may@v0.0.4
```

## Intro

May is a golang SDK for Meilisearch

## Usage

### Declare data Model with tag

```go
type Sample struct {
SampleID string `json:"sampleID" may:"index"`
Title    string `json:"title" may:"1"`
Des      string `json:"des" may:"2"`
NoSearch string `json:"noSearch" may:"-"`
Test     string `json:"test"`
}
```

- **index**: primary key field
- **integer**: searchable field, weight rank

### Init Client And Index

```go
may.InitClient("HTTP://MEILI.SEARCH.ADDR:7700", "YOUR-KEY")
err := may.CreateIndex(Sample{})
if err != nil {
    return err
}
err = may.Add(Sample{
    SampleID: "1",
    Title:    "test1",
    Des:      "test1",
    Test:     "okok",
})
    return err
```

### Search

search with keyword and pass pointer to get value

```go
var s []Sample
err := may.NewFinder().Search("test", &s)
if err != nil {
    t.Fatal(err)
}
```
