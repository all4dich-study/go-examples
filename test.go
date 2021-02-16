package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"github.com/all4dich/golang/buildanalysis/oebuildjobs"
	"io/ioutil"
	"os"
)

func test() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State, Country string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	v := Result{Name: "none", Phone: "none"}

	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
			<Country>Wacanda</Country>
			<Job>Engineer</Job>
		</Person>
	`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)
	fmt.Printf("Country: %v\n", v.Country)
}

func main() {

	type Account struct {
		Name     string `xml:"name"`
		Email    string `xml:"email"`
		Username string `xml:"username"`
	}

	type Provider struct {
		Name    string `xml:"name"`
		Host    string `xml:"host"`
		Port    int    `xml:"port"`
		Scheme  string `xml:"scheme"`
		Url     string `xml:"url"`
		Version string `xml:"version"`
	}

	type GerritChange struct {
		Project       string  `xml:"project"`
		Branch        string  `xml:"branch"`
		Id            string  `xml:"id"`
		Number        int     `xml:"number"`
		Subject       string  `xml:"subject"`
		CommitMessage string  `xml:"commitMessage"`
		Owner         Account `xml:"owner"`
		Url           string  `xml:"url"`
	}

	type PatchSet struct {
		Number    int     `xml:"number"`
		Revision  string  `xml:"revision"`
		Ref       string  `xml:"ref"`
		Draft     bool    `xml:"draft"`
		Uploader  Account `xml:"uploader"`
		Author    Account `xml:"Author"`
		Parents   string  `xml:"parents>string"`
		CreatedOn string  `xml:"createdOn"`
	}

	type Files []string

	type Approval struct {
		Type  string `xml:"type"`
		Value string `xml:"value"`
	}

	type TEvent struct {
		Provider     Provider     `xml:"provider"`
		Account      Account      `xml:"account"`
		GerritChange GerritChange `xml:"change"`
		PatchSet     PatchSet     `xml:"patchSet"`
		Files        Files        `xml:"files>string"`
		Comment      string       `xml:"comment"`
		Approvals    []Approval   `xml:"approvals>com.sonyericsson.hudson.plugins.gerrit.gerritevents.dto.attr.Approval"`
	}

	type TriggeredItem struct {
		BuildNumber int    `xml:"buildNumber"`
		ProjectId   string `xml:"projectId"`
	}

	type RetriggerContext struct {
		ThisBuild TriggeredItem   `xml:"thisBuild"`
		Others    []TriggeredItem `xml:"others>triggeredItemEntity"`
	}

	type RetriggerEvent struct {
		Context RetriggerContext `xml:"context"`
	}

	type EachParameter struct {
		Name  string `xml:"name"`
		Value string `xml:"value"`
	}
	type Build struct {
		XMLName       xml.Name        `xml:"build"`
		StartTime     int             `xml:"startTime"`
		Duration      int             `xml:"duration"`
		Result        string          `xml:"result"`
		Host          string          `xml:"builtOn"`
		BuildEvent    TEvent          `xml:"actions>com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.BadgeAction>tEvent"`
		RetriggerInfo RetriggerEvent  `xml:"actions>com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.actions.RetriggerAction"`
		Parameters    []EachParameter `xml:"actions>hudson.model.ParametersAction>parameters>hudson.model.StringParameterValue"`
		Description   string          `xml:"description"`
		/*
			Provider     Provider     `xml:"actions>com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.BadgeAction>tEvent>provider"`
			Account      Account      `xml:"actions>com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.BadgeAction>tEvent>account"`
			GerritChange GerritChange `xml:"actions>com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.BadgeAction>tEvent>change"`
			PatchSet     PatchSet     `xml:"actions>com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.BadgeAction>tEvent>patchSet"`
			Files        Files        `xml:"actions>com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.BadgeAction>tEvent>files>string"`
			Comment      string       `xml:"actions>com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.BadgeAction>tEvent>comment"`
			Approvals    []Approval   `xml:"actions>com.sonyericsson.hudson.plugins.gerrit.trigger.hudsontrigger.BadgeAction>tEvent>approvals>com.sonyericsson.hudson.plugins.gerrit.gerritevents.dto.attr.Approval"`
		*/
	}
	//v := Build{}
	v := oebuildjobs.VerifyBuild{}

	file, err := os.Open("/Users/sunjoo/temp/jenkins_home/jobs/starfish-drd4tv-verify-h15/builds/10001/build.xml")
	if err != nil {
		fmt.Println("Error: Cannot open a file")
		os.Exit(1)
	}
	r := bufio.NewReader(file)
	dat, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("Error: Cannot read a string")
		os.Exit(1)
	}
	err = xml.Unmarshal(dat, &v)
	fmt.Println(v.StartTime)
	fmt.Println(v.Duration)
	fmt.Println(v.Result)
	fmt.Println(v.Host)
	for _, each_parameter := range v.Parameters {
		fmt.Println(each_parameter.Name, each_parameter.Value)
	}
	fmt.Println(v.Description)

}
