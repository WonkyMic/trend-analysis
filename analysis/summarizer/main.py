from trend_db import TrendDB
from summarize import Summarizer

def main():

    db = TrendDB()
    db.connect()
    empty_summaries = db.get_empty_article_summaries()
    db.close()
    
    # Initialize the summarizer
    summarizer = Summarizer()
    
    # Summarize each article and update the database
    for article in empty_summaries:
        summary = summarizer.summarize(article[1])
        db = TrendDB()
        db.connect()
        db.update_article_summary(article[0], summary)
        db.close()
        print(f"Summarized article {article[0]}")

    print("-- Summarization complete --")


if __name__ == "__main__":
    main()