package main

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"fmt"
	"os"
	"strings"
	"time"
	"os/exec"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
)

var choices = []string{"Display Current Julian Date", "Convert Gregorian Date to Julian Date", "Convert Julian Date to Gregorian Date", "Exit"}

type model struct {
	cursor int
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("\n")
	s.WriteString("CelestiWatch - Galactic Chronometer\n\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("[=] ")
		} else {
			s.WriteString("[ ] ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	
	s.WriteString("\n(press q to quit)\n\n")
	s.WriteString("---------------\n\n")

	return s.String()
}

func main() {
	p := tea.NewProgram(model{})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
			clearScreen()
			switch(m.choice) {

				case "Display Current Julian Date" : 
					fmt.Printf("--------------------\n\n")
					fmt.Printf("[=] Display Current Julian Date\n")
					fmt.Printf("# Peek into the Julian side of the time vortex and see today's date like it's the star of its own cosmic show!\n\n")
					// for i := 0; i < 10; i++ { 
						julianDate := timeToJulianDate(time.Now())
						fmt.Printf("\rCurrent Julian Date : %.6f", julianDate)
						time.Sleep(time.Second)
					// }
					fmt.Printf("\n\n")

				case "Convert Gregorian Date to Julian Date" : 
					fmt.Printf("--------------------\n\n")
					fmt.Printf("[=] Convert Gregorian Date to Julian Date\n")
					fmt.Printf("# Transform your everyday Gregorian date into a swanky, retro Julian date – because time travel should be a fashion statement!\n\n")
					fmt.Print("Enter Gregorian date (YYYY-MM-DD): \n")
					var input string
					fmt.Scanln(&input)
					fmt.Print("\n")
					// Parse the user input as a Gregorian date
					gregorianDate, err := time.Parse("2006-01-02", input)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}
			
					// Convert the Gregorian date to Julian date
					julianDate := gregorianToJulian(gregorianDate)
			
					// Print the result
					fmt.Printf("Inputted Gregorian Date: %s\n", gregorianDate.Format("2006-01-02"))
					fmt.Printf("Converted Julian Date: %.6f\n", julianDate)
					fmt.Print("\n")

				case "Convert Julian Date to Gregorian Date" : 
					fmt.Printf("--------------------\n\n")
					fmt.Printf("[=] Convert Julian Date to Gregorian Date\n")
					fmt.Printf("# Unleash your inner time wizard and translate a Julian date into a Gregorian one – it's like turning ancient scrolls into a modern-day calendar dance!\n\n")
					fmt.Print("Enter Julian date : \n")
					var input float64
					fmt.Scanln(&input)
					fmt.Print("\n")
			
					// Convert the Gregorian date to Julian date
					gregorianDate := julianToGregorian(input)
			
					// Print the result
					// Print the result
					fmt.Printf("Inputted Julian Date: %.6f\n", input)
					fmt.Printf("Converted Gregorian Date: %s\n", gregorianDate.Format("2006-01-02"))
					fmt.Print("\n")

				case "Exit" : 
					fmt.Printf("--------------------\n\n")
					fmt.Println("Terminated.")
					os.Exit(0)
			}
	}
}

func timeToJulianDate(t time.Time) float64 {
	julianEpoch := 2451545.0

	// Convert the given time to Julian Date
	days := t.Sub(time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)).Hours() / 24.0
	julianDate := julianEpoch + days

	return julianDate
}

func gregorianToJulian(t time.Time) float64 {
	// Constants for Julian Date calculation
	julianEpoch := 2451545.0

	// Convert the given time to Julian Date
	days := t.Sub(time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)).Hours() / 24.0
	julianDate := julianEpoch + days

	return julianDate
}

func julianToGregorian(julianDate float64) time.Time {
	// Constants for Julian Date calculation
	julianEpoch := 2451545.0

	// Calculate the number of days since the Julian epoch
	daysSinceEpoch := julianDate - julianEpoch

	// Convert daysSinceEpoch to a duration
	daysDuration := time.Duration(daysSinceEpoch * float64(time.Hour*24))

	// Calculate the Gregorian date
	gregorianDate := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC).Add(daysDuration)

	return gregorianDate
}

func clearScreen() {
	// Clear the console screen based on the operating system
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}