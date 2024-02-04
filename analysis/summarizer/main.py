from trend_db import TrendDB
from summarize import Summarizer
from concurrent.futures import ProcessPoolExecutor

def summarize_and_update(article, summarizer):
    summary = summarizer.summarize(article[1])
    db = TrendDB()
    db.connect()
    db.update_article_summary(article[0], summary)
    db.close()
    
def main():

    db = TrendDB()
    db.connect()
    empty_summaries = db.get_empty_article_summaries()
    db.close()
    
    # Initialize the summarizer
    summarizer = Summarizer()
    
    # Summarize each article and update the database
    with ProcessPoolExecutor() as executor:
        for article in empty_summaries:
            executor.submit(summarize_and_update, article, summarizer)

    print("-- Summarization complete --")

if __name__ == "__main__":
    main()