# trend-analysis
Using Wikipedia's Pageviews API we are able to aggregate trending topics and extract the contents of the article.

# TODO
- [ ] Save the aggregation into a persistance store
- [ ] Compare topics over time to see emerging ideas
- [ ] Pass the article extracts to an LLM to summarize data
- [ ] Identify significant page edits within trending articles
- [ ] Convert to CLI or batch process

# Developer notes
To execute
```go run .```

The Pageviews API returns 1000 records by default. The constant `breakpoint` within the main class serves two functions. (1) It prevents calling the extract API for all records and (2) limits the number of records printed to output.