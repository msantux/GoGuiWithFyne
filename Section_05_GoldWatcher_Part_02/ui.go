package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	// get current price of gold
	openPrice, currentPrice, priceChange := app.getPriceText()

	// put price information into a container
	priceContent := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange,
	)
	app.PriceContainer = priceContent

	// get toolbar
	toolbar := app.getToolBar()
	app.Toolbar = toolbar

	// get app tabs
	priceTabContent := app.pricesTab()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), canvas.NewText("Holdings content goes here", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to the window
	finalContent := container.NewVBox(priceContent, toolbar, tabs)
	app.MainWindow.SetContent(finalContent)

	go func() {
		for range time.Tick(time.Second * 5) {
			app.refreshPriceContent()
		}
	}()
}

func (app *Config) refreshPriceContent() {
	app.InfoLog.Println("refreshing prices")

	open, current, change := app.getPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceChartContainer.Refresh()
}
