package main

import (
	"GoGuiWithFyne/Section_06_GoldWatcher_Part_03/repository"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) holdingsTab() *fyne.Container {
	app.HoldingsData = app.getHoldingsSlice()
	app.HoldingsTable = app.getHoldingsTable()

	holdingsContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, app.HoldingsTable),
	)

	return holdingsContainer
}

func (app *Config) getHoldingsTable() *widget.Table {
	t := widget.NewTable(
		func() (int, int) {
			return len(app.HoldingsData), len(app.HoldingsData[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))

			return ctr
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == (len(app.HoldingsData[0])-1) && i.Row != 0 {
				// last cell - put in a button
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "", func(delete bool) {
						if delete {
							id, _ := strconv.Atoi(app.HoldingsData[i.Row][0].(string))
							err := app.DB.DeleteHolding(int64(id))
							if err != nil {
								app.ErrorLog.Println(err)
							}
						}
						// refresh holdings table
						app.refreshHoldingsTableContent()
					}, app.MainWindow)
				})

				w.Importance = widget.HighImportance
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					w,
				}
			} else {
				// we just puting textual information
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(app.HoldingsData[i.Row][i.Col].(string)),
				}
			}
		})

	colWidths := []float32{50, 200, 200, 200, 110}
	for i := 0; i < len(colWidths); i++ {
		t.SetColumnWidth(i, colWidths[i])
	}
	return t
}

func (app *Config) getHoldingsSlice() [][]interface{} {
	var slice [][]interface{}

	holdings, err := app.currentHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	slice = append(slice, []interface{}{"ID", "Amount", "Price", "Date", "Delete?"})
	for _, x := range holdings {
		var currentRow []interface{}

		currentRow = append(currentRow, strconv.FormatInt(x.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%d toz", x.Amount))
		currentAmount := float32(x.PurchasePrice)
		currentAmount = currentAmount / 100
		currentRow = append(currentRow, fmt.Sprintf("$%.2f", currentAmount))
		currentRow = append(currentRow, x.PurchaseDate.Format("2006-01-02"))
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}

	return slice
}

func (app *Config) currentHoldings() ([]repository.Holdings, error) {
	holdings, err := app.DB.AllHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	return holdings, nil
}
