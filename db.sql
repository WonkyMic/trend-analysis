CREATE TABLE public.page_views (
	title varchar NULL,
	"date" date NULL,
	"views" int NULL,
	CONSTRAINT page_views_unique UNIQUE (title,"date")
);

CREATE TABLE public.article_summary (
	title varchar NULL,
	"extract" varchar NULL,
	summary varchar NULL,
	last_summarize_date date NULL,
	CONSTRAINT article_summary_unique UNIQUE (title)
);

