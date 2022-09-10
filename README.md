<p align="center" width="100%">
    <img width="45%" src="./ui/src/assets/svg.svg"> 
</p>

# Atlas | [Releases](https://github.com/knexguy101/Atlas/releases)
A full fledged Nike Browser bot. Currently only supports the US region but other regions are coming soon (feel free to open a PR!). Please be aware of our licensing before working with our source in a private environment (GNU GPL3). 
</br>
</br>
Atlas runs on [Vue](https://vuejs.org/), [Lorca](https://github.com/zserge/lorca), and [Playwright](https://github.com/playwright-community/playwright-go)!

## Features
- Fully automated US SNKRS entry
- Automatically add address/payment to profiles
- Multi task support
- Multi account/proxy support
- Timer start or Release Monitoring
- GUI

## Building from scratch
```markdown
# Downloading the files
git fetch https://github.com/knexguy101/Atlas
cd Atlas

# Build the bot client
go build .

# Build the UI
cd ui
npm install
npm run build
```
