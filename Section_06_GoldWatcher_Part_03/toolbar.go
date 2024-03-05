package main

import (
	"GoGuiWithFyne/Section_06_GoldWatcher_Part_03/repository"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) getToolBar() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			app.addHoldingsDialog()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			app.refreshPriceContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			w := app.showPreferences()
			w.Resize(fyne.NewSize(300, 200))
			w.Show()
		}),
	)

	return toolbar
}

func (app *Config) addHoldingsDialog() dialog.Dialog {
	amountEntry := widget.NewEntry()
	purchaseDateEntry := widget.NewEntry()
	purchasePriceEntry := widget.NewEntry()

	app.AddHoldingsPurchaseAmountEntry = amountEntry
	app.AddHoldingsPurchaseDateEntry = purchaseDateEntry
	app.AddHoldingsPurchasePriceEntry = purchasePriceEntry

	dateValidator := func(s string) error {
		if _, err := time.Parse("2006-01-02", s); err != nil {
			return err
		}

		return nil
	}
	purchaseDateEntry.PlaceHolder = "YYYY-MM-DD"
	purchaseDateEntry.Validator = dateValidator

	isIntValidator := func(s string) error {
		if _, err := strconv.Atoi(s); err != nil {
			return err
		}

		return nil
	}
	amountEntry.Validator = isIntValidator

	isFloatValidator := func(s string) error {
		if _, err := strconv.ParseFloat(s, 32); err != nil {
			return err
		}

		return nil
	}
	purchasePriceEntry.Validator = isFloatValidator

	// create dialog
	addForm := dialog.NewForm(
		"Add Gold",
		"Add",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Amount in toz:", Widget: amountEntry},
			{Text: "Purchase Price:", Widget: purchasePriceEntry},
			{Text: "Purchase date:", Widget: purchaseDateEntry},
		},
		func(valid bool) {
			if valid {
				amount, _ := strconv.Atoi(amountEntry.Text)
				price, _ := strconv.ParseFloat(purchasePriceEntry.Text, 32)
				price = price * 100
				date, _ := time.Parse("2006-01-02", purchaseDateEntry.Text)

				_, err := app.DB.InsertHolding(repository.Holdings{
					Amount:        amount,
					PurchasePrice: int(price),
					PurchaseDate:  date,
				})
				if err != nil {
					app.ErrorLog.Println(err)
				}
				app.refreshHoldingsTableContent()
			}
		},
		app.MainWindow)

	// size and show dialog
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}

func (app *Config) showPreferences() fyne.Window {
	win := app.App.NewWindow("Preferences")

	lbl := widget.NewLabel("Prefered currency")
	cur := widget.NewSelect([]string{"BRL", "CAD", "GBP", "USD"}, func(s string) {
		currency = s
		app.App.Preferences().SetString("currency", currency)
	})
	cur.Selected = currency

	btn := widget.NewButton("Save", func() {
		win.Close()
		app.refreshPriceContent()
	})
	btn.Importance = widget.HighImportance

	win.SetContent(container.NewVBox(lbl, cur, btn))

	return win
}
