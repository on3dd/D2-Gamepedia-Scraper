# D2-Gamepedia-Scraper
Grabs the heroes responses from [dota2.gamepedia.com](https://dota2.gamepedia.com) and outputs them in a browser
![Preview](https://i.imgur.com/Hl6dKc7.png)
# Usage
## Install
Make sure Go is already installed on your PC.

Clone this repository and install all required dependencies.

## Start 
Type the following code in the terminal:
```
go run main.go
```

Then open a browser and go to `http://localhost:8080/`.

## Get results
Type `http://localhost:8080/heroes/{Hero}` in browsers adress bar, where *{Hero}* - some hero name. 

## TODO:
- [ ] Fix the bug with Lina and Crystal Maiden audio playback
- [ ] Add an exception for pages with the broken audios
