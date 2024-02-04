from transformers import pipeline

class Summarizer:
    def __init__(self):
        # Initialize the summarizer
        self.summarizer = pipeline("summarization")

    def summarize(self, article_text):
        # Split the article into chunks of 1024 tokens
        chunks = [article_text[i:i+1024] for i in range(0, len(article_text), 1024)]

        # Summarize each chunk and concatenate the summaries
        summaries = [self.summarizer(chunk, max_length=20, min_length=10, do_sample=False)[0]['summary_text'] for chunk in chunks]
        summary = " ".join(summaries)
        return summary