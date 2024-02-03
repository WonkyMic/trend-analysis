import psycopg2
from dotenv import load_dotenv
import os

class TrendDB:
    def __init__(self):
        load_dotenv()
        self.db_name = os.getenv('DB_NAME')
        self.user = os.getenv('DB_USER')
        self.password = os.getenv('DB_PASS')
        self.host = os.getenv('DB_HOST')
        self.port = os.getenv('DB_PORT')

    def connect(self):
        self.conn = psycopg2.connect(
            dbname=self.db_name,
            user=self.user,
            password=self.password,
            host=self.host,
            port=self.port
        )

    def close(self):
        self.conn.close()

    def get_page_views(self):
        cursor = self.conn.cursor()
        cursor.execute("SELECT * FROM public.page_views")
        trends = cursor.fetchall()
        cursor.close()
        return trends
    
    def get_article_summary(self, title):
        cursor = self.conn.cursor()
        cursor.execute("SELECT * FROM public.article_summary WHERE title = %s", (title,))
        summary = cursor.fetchone()
        cursor.close()
        return summary
    
    def get_empty_article_summaries(self):
        cursor = self.conn.cursor()
        cursor.execute("SELECT title, extract FROM public.article_summary WHERE summary IS NULL")
        summaries = cursor.fetchall()
        cursor.close()
        return summaries
    
    def get_article_summaries_last_updated_before(self, date):
        cursor = self.conn.cursor()
        cursor.execute("SELECT * FROM public.article_summary WHERE last_summarize_date < %s", (date,))
        summaries = cursor.fetchall()
        cursor.close()
        return summaries
    
    def update_article_summary(self, title, summary):
        cursor = self.conn.cursor()
        cursor.execute("UPDATE public.article_summary SET summary = %s, last_summarize_date = CURRENT_TIMESTAMP WHERE title = %s", (summary, title))
        self.conn.commit()
        cursor.close()

    def insert_trend(self, trend):
        cursor = self.conn.cursor()
        cursor.execute("INSERT INTO trends (name, volume) VALUES (%s, %s)", (trend['name'], trend['volume']))
        self.conn.commit()
        cursor.close()