# LDS-Talks-Scraper
Project to scrape all LDS talks from general conference, potentially analyze them and map to scripture citations. Probably going into a search functionality in the long term

#### Talk processing

To process the talks, `go run parallel.go <start_position> <end_position> <n_threads>` where `<start_position>` and `<end_position>` are the numbers of the talks in the url that go here `https://scriptures.byu.edu/content/talks_ajax/talk_num`. You can run this to add talks as new ones are added at each General Conference. 