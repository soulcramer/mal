package mal

// Anime ...
type Anime struct {
	SeriesAnimedbID   string        `json:"series_animedb_id" xml:"series_animedb_id"`
	SeriesTitle       string        `json:"series_title" xml:"series_title"`
	SeriesSynonyms    string        `json:"series_synonyms" xml:"series_synonyms"`
	SeriesType        string        `json:"series_type" xml:"series_type"`
	SeriesEpisodes    string        `json:"series_episodes" xml:"series_episodes"`
	SeriesStatus      string        `json:"series_status" xml:"series_status"`
	SeriesStart       string        `json:"series_start" xml:"series_start"`
	SeriesEnd         string        `json:"series_end" xml:"series_end"`
	SeriesImage       string        `json:"series_image" xml:"series_image"`
	MyID              string        `json:"my_id" xml:"my_id"`
	MyWatchedEpisodes string        `json:"my_watched_episodes" xml:"my_watched_episodes"`
	MyStartDate       string        `json:"my_start_date" xml:"my_start_date"`
	MyFinishDate      string        `json:"my_finish_date" xml:"my_finish_date"`
	MyScore           string        `json:"my_score" xml:"my_score"`
	MyStatus          string        `json:"my_status" xml:"my_status"`
	MyRewatching      []interface{} `json:"my_rewatching" xml:"my_rewatching"`
	MyRewatchingEp    string        `json:"my_rewatching_ep" xml:"my_rewatching_ep"`
	MyLastUpdated     string        `json:"my_last_updated" xml:"my_last_updated"`
	MyTags            []interface{} `json:"my_tags" xml:"my_tags"`
}
