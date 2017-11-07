package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"baliance.com/gooxml/document"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

// Employee struct for the literal information
type Employee struct {
	Name           string
	Company        string
	Position       string
	YrsOfStay      string
	DirectorName   string
	ContactNumber  string
	EmailAddress   string
	Address        string
	CompanyAddress string
	EffectiveDate  string
	// Choices
	Resign     string
	LooksGood  string
	ReallySure string
}

// Init function will initialize to get the user's decision
func (e Employee) Init() Employee {
	// Some greetings
	fmt.Print("Thanks you for trusting RESIGNATION LETTER GENERATOR. Your number 1 tool for your graceful resignation.\n\n")
	fmt.Print("By using this tool - it means that you're now to ready to get out from your comfort zone, to level up your career, to earn more than your current salary, or perhaps you're just bored to your current toxic office life.\n\n")

	fmt.Print("Do you wish to continue? [y/n]: ")
	e.Resign, _ = reader.ReadString('\n')
	e.Resign = strings.Replace(e.Resign, "\n", "", -1)
	return e
}

// GetInformation func will collect the information needed for the template
func (e Employee) GetInformation() Employee {

	fmt.Print("\n\nPlease Provide Your Personal Information.\n\n")

	fmt.Print("Your Fullname: ")
	e.Name, _ = reader.ReadString('\n')

	fmt.Print("Your Current Address: ")
	e.Address, _ = reader.ReadString('\n')

	fmt.Print("Your Contact #: ")
	e.ContactNumber, _ = reader.ReadString('\n')

	fmt.Print("Your Email Address: ")
	e.EmailAddress, _ = reader.ReadString('\n')

	fmt.Print("Company's Name / Employer's Name: ")
	e.Company, _ = reader.ReadString('\n')

	fmt.Print("Office's Address / Location: ")
	e.CompanyAddress, _ = reader.ReadString('\n')

	fmt.Print("Your Current Position: ")
	e.Position, _ = reader.ReadString('\n')

	fmt.Print("No. of Years of Stay: ")
	e.YrsOfStay, _ = reader.ReadString('\n')

	fmt.Print("Your Director's Name (Including Salutations): ")
	e.DirectorName, _ = reader.ReadString('\n')

	fmt.Print("Your Resignation's Effectivity Date: ")
	e.EffectiveDate, _ = reader.ReadString('\n')

	e = e.CleanInfo()

	e.Validate()

	return e
}

// Validate func will let the user to validate his/her info
func (e Employee) Validate() {
	fmt.Print("\n\n")
	fmt.Print("Preview:\n\n")
	fmt.Print("Fullname: ", e.Name+"\n")
	fmt.Print("Current Address: ", e.Address+"\n")
	fmt.Print("Contact #: ", e.ContactNumber+"\n")
	fmt.Print("Email Address: ", e.EmailAddress+"\n")
	fmt.Print("Company: ", e.Company+"\n")
	fmt.Print("Office Address: ", e.CompanyAddress+"\n")
	fmt.Print("Position: ", e.Position+"\n")
	fmt.Print("Years of Stay: ", e.YrsOfStay+"\n")
	fmt.Print("Director's Name (Including Salutations): ", e.DirectorName+"\n")
	fmt.Print("Effectivity Date: ", e.EffectiveDate+"\n")
	fmt.Print("\n\n")
	fmt.Printf("Looks good? [y/n]: ")
	e.LooksGood, _ = reader.ReadString('\n')
	e.LooksGood = strings.Replace(e.LooksGood, "\n", "", -1)

	if e.LooksGood == "y" || e.LooksGood == "Y" {
		fmt.Print("\n")
		fmt.Print("Are you really sure? [y/n]: ")
		e.ReallySure, _ = reader.ReadString('\n')
		e.ReallySure = strings.Replace(e.ReallySure, "\n", "", -1)
		if e.ReallySure == "y" || e.ReallySure == "Y" {
			fmt.Print("\n")
			fmt.Print("Generating resignation letter...\n")
			e.GenerateLetter()
			fmt.Print("Done on generating your Resignation Letter. See the Resignation_Letter.docx file.\n\n")
		} else {
			fmt.Print("Fuck you! Such a waste of time!")
		}
	} else {
		e.GetInformation()
	}
}

// CleanInfo func will remove the breakline on the information
func (e Employee) CleanInfo() Employee {
	e.Name = strings.Replace(e.Name, "\n", "", -1)
	e.Company = strings.Replace(e.Company, "\n", "", -1)
	e.Position = strings.Replace(e.Position, "\n", "", -1)
	e.DirectorName = strings.Replace(e.DirectorName, "\n", "", -1)
	e.ContactNumber = strings.Replace(e.ContactNumber, "\n", "", -1)
	e.EmailAddress = strings.Replace(e.EmailAddress, "\n", "", -1)
	e.Address = strings.Replace(e.Address, "\n", "", -1)
	e.EffectiveDate = strings.Replace(e.EffectiveDate, "\n", "", -1)
	e.YrsOfStay = strings.Replace(e.YrsOfStay, "\n", "", -1)
	e.Company = strings.Replace(e.Company, "\n", "", -1)
	e.CompanyAddress = strings.Replace(e.CompanyAddress, "\n", "", -1)

	return e
}

// Decline func shows a bad conclusion
func (e Employee) Decline() {
	fmt.Print("K fine.")
}

// GenerateLetter func will generate instant resignation letter
func (e Employee) GenerateLetter() {
	doc := document.New()

	para := doc.AddParagraph()
	run := para.AddRun()
	run.Properties().SetSize(12)

	run.AddText(e.Name)
	run.AddBreak()
	run.AddText(e.Address)
	run.AddBreak()
	run.AddText(e.ContactNumber)
	run.AddBreak()
	run.AddText(e.EmailAddress)

	run.AddBreak()
	run.AddBreak()
	run.AddText(e.EffectiveDate)

	run.AddBreak()
	run.AddBreak()
	run.AddText(e.DirectorName)
	run.AddBreak()
	run.AddText("Development Manager")
	run.AddBreak()
	run.AddText(e.Company)
	run.AddBreak()
	run.AddText(e.CompanyAddress)
	run.AddBreak()

	run.AddBreak()
	run.AddText("Dear " + e.DirectorName + ":")
	run.AddBreak()

	para1 := doc.AddParagraph()
	run = para1.AddRun()
	run.Properties().SetSize(12)
	run.AddTab()
	run.AddText(fmt.Sprintf("Please accept this letter as formal notification that I am leaving my position as %v with %v.", e.Position, e.Company))
	run.AddBreak()

	para2 := doc.AddParagraph()
	run = para2.AddRun()
	run.Properties().SetSize(12)
	run.AddTab()
	run.AddText(fmt.Sprintf("I am grateful for all the opportunities this company has provided me for more than %v years, including the training, experiences on software development methodologies, programming technologies, and most especially the fun and great relationship I had with my colleagues.", e.YrsOfStay))
	run.AddBreak()

	para3 := doc.AddParagraph()
	run = para3.AddRun()
	run.Properties().SetSize(12)
	run.AddTab()
	run.AddText("I'll do my best to wrap up my remaining tasks and aid in the transition to those who will assume the tasks I am allocated to.")
	run.AddBreak()

	para4 := doc.AddParagraph()
	run = para4.AddRun()
	run.Properties().SetSize(12)
	run.AddTab()
	run.AddText("Thank you. Wishing more power and success to this company.")
	run.AddBreak()

	para5 := doc.AddParagraph()
	run = para5.AddRun()
	run.Properties().SetSize(12)
	run.AddBreak()
	run.AddText("Sincerely,")
	run.AddBreak()
	run.AddText(e.Name)

	doc.SaveToFile("Resignation_Letter.docx")
}
