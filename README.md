# D2-Gamepedia-Scraper
Grabs the heroes responses from dota2.gamepedia.com and outputs them in a browser
![Preview](https://i.imgur.com/Hl6dKc7.png)
# Usage
## Start 
Type the following code in the terminal:
```
go run main.go
```

## Get results
Type `http://localhost:8080/heroes/{Hero}` in browsers adress bar, where *{Hero}* - some hero name. 

**Note that hero name must start with a capital letter, and all the whitespaces should be replaced by underscores.**

## TODO's:
- [ ] Fix the bug with Lina and Crystal Maiden audio playback
- [ ] Add a 404 page
- [ ] Add an exception for pages with the broken audios
