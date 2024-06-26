package main

import (
	b "aio/auth"
	"aio/handlers"
	"aio/helpers"
	"aio/logging"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	proxies []string
	tokens  []string
	auths   []string
	err     error
	module  string
	tweets  []string
)
var clear map[string]func()

// * KeyAuth Application Details *//
var name = "TwitterAIO"
var ownerid = "aHUUIOGNKw"
var version = "1.0"

// * API SET UP VALUES ^^^^ *//
var username = "" //* Keep Clear
var password = "" //* Keep Clear
var key = ""      //* Keep Clear
func ClearConsole() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func animateBanner() {
	file, err := os.Open("input/banner.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	c := color.New(color.FgHiCyan)

	for _, line := range lines {
		c.Print(line)
		time.Sleep(200 * time.Millisecond)
		fmt.Print("\n")
	}

	// Pause for a short duration after printing the banner
	time.Sleep(1 * time.Second)
}
func Authenticator() {
	b.Api(name, ownerid, version) // Important to set up the API Details

	reader := bufio.NewReader(os.Stdin)

	ClearConsole()

	// fmt.Println("\n\n Connecting..")
	b.Init()

	// fmt.Println("\n App Data:")
	// fmt.Println(" Number of users:", b.NumUsers)
	// fmt.Println(" Number of online users:", b.NumOnlineUsers)
	// fmt.Println(" Number of keys:", b.NumKeys)
	// fmt.Println(" Application Version:", version)
	// fmt.Println(" Customer Panel Link:", b.CustomerPanelLink)

	fmt.Println("\n [1] Login\n [2] Register\n [3] Upgrade\n [4] License key only\n\n Choose option: ")

	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	switch char {
	case '1':
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSuffix(username, "\n")

		fmt.Println("Enter Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSuffix(password, "\n")

		b.Login(username, password)

		break
	case '2':
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSuffix(username, "\n")

		fmt.Println("Enter Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSuffix(password, "\n")

		fmt.Println("Enter License: ")
		key, _ := reader.ReadString('\n')
		key = strings.TrimSuffix(key, "\n")

		b.Register(username, password, key)

		break
	case '3':
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSuffix(username, "\n")

		fmt.Println("Enter License: ")
		key, _ := reader.ReadString('\n')
		key = strings.TrimSuffix(key, "\n")

		b.Upgrade(username, key)

		break
	case '4':
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter License: ")
		key, _ := reader.ReadString('\n')
		key = strings.TrimSuffix(key, "\n")

		b.License(key)

		break
	default:
		fmt.Println("Invalid Selection")
		os.Exit(0)
		break
	}

	// fmt.Println("\n User data:")
	// fmt.Println(" Username:", b.Username)
	// fmt.Println(" IP address:", b.Ip)
	// fmt.Println(" Hardware-Id:", b.Hwid)
	// fmt.Println(" Created at:", b.Createdate)
	// fmt.Println(" Last login at:", b.Lastlogin)
	// fmt.Println(" Subscription:", b.Subscription)

	/* --> Extra Functions <--
	* User Variables *
	b.SetVar("VariableName", "VariableData") // Set up User Variable
	b.GetVar("VariableName") // Get User Variable

	* Get Public Variables * - https://keyauth.cc/dashboard/app/variables/

	b.Var("VariableName") // Get Public Variable

	Example:
	var publicVariable = b.Var("VariableName")
	fmt.Println("Variable Content: " + publicVariable)

	* Webhooks * - https://keyauth.cc/dashboard/app/webhooks/
	b.Webhook("WebhookName", "WebhookData") // Send Webhook

	Example:

	var WbData = b.Webhook("Webhook ID", "?type=test")
	fmt.Println("Webhook Data: " + WbData)

	* Logs * - https://keyauth.cc/dashboard/app/settings/
	b.Log("Message") // Send Log to Webhook of your choice ^^
	*/

}
func printBanner() {
	// Set colors for different sections
	sectionHeaderColor := color.New(color.FgCyan, color.Bold)
	menuOptionColor := color.New(color.FgHiYellow, color.Bold)

	// Print the banner
	sectionHeaderColor.Println("[Account Management]")
	menuOptionColor.Println("[1] Create Account DB [Converter]")
	menuOptionColor.Println("[2] Token Checker [Filter Bad/Login-Failed/Locked/Suspended Accounts]\n")

	sectionHeaderColor.Println("[Mass]")
	menuOptionColor.Println("[3] Mass Tweet")
	menuOptionColor.Println("[4] Mass Like")
	menuOptionColor.Println("[5] Mass Retweet")
	menuOptionColor.Println("[6] AI Comments [GPT3.5]\n")
}

func LoadTxts() {
	tokens, err = helpers.ReadLinesFromFile("input/tokens.txt")
	if err != nil {
		logging.Log(logging.Error, err.Error())
	}
	if len(tokens) == 0 {
		logging.Log(logging.Warning, "0 Tokens in File")
	}
	proxies, err = helpers.ReadLinesFromFile("input/proxies.txt")
	if err != nil {
		logging.Log(logging.Error, err.Error())
	}
	if len(proxies) == 0 {
		logging.Log(logging.Warning, "0 Proxies in File")
	}
	auths, err = helpers.ReadLinesFromFile("input/auth.txt")
	if err != nil {
		logging.Log(logging.Error, err.Error())
	}
	if len(auths) == 0 {
		logging.Log(logging.Warning, "0 Auths in File")
	}
	tweets, err = helpers.ReadLinesFromFile("input/tweets/tweets.txt")
	if err != nil {
		logging.Log(logging.Error, err.Error())
	}
	if len(tweets) == 0 {
		logging.Log(logging.Warning, "0 Mobile in File")
	}

}
func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	// Authenticator()
}
func main() {
restart:
	animateBanner()
	fmt.Println()
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	printBanner()
	LoadTxts()
	logging.InitLogger(logging.Info)

	fmt.Print("Module: ")
	fmt.Scanln(&module)

	switch module {
	case "1":
		var threadsstr string
		fmt.Print("Threads: ")
		fmt.Scanln(&threadsstr)
		threadsstr = strings.TrimSpace(threadsstr)
		threadsint, _ := strconv.Atoi(threadsstr)

		handlers.ConverterHandler(auths, proxies, threadsint)
		goto restart
	case "2":
		var threadsstr string
		fmt.Print("Threads: ")
		fmt.Scanln(&threadsstr)
		threadsstr = strings.TrimSpace(threadsstr)
		threadsint, _ := strconv.Atoi(threadsstr)
		handlers.CheckerHandler(tokens, threadsint)
		goto restart
	case "3":
		var threadsstr string
		fmt.Print("Threads: ")
		fmt.Scanln(&threadsstr)
		threadsstr = strings.TrimSpace(threadsstr)
		threadsint, _ := strconv.Atoi(threadsstr)
		handlers.HandleTweets(tokens, tweets, threadsint)
		goto restart
	case "4":
		var tweet_id string
		var threadsstr string
		fmt.Print("Threads: ")
		fmt.Scanln(&threadsstr)
		threadsstr = strings.TrimSpace(threadsstr)
		threadsint, _ := strconv.Atoi(threadsstr)
		fmt.Print("TweetID: ")
		fmt.Scanln(&tweet_id)
		handlers.HandleLike(tokens, tweet_id, threadsint)
		goto restart
	case "5":
		var tweet_id string
		var threadsstr string
		fmt.Print("Threads: ")
		fmt.Scanln(&threadsstr)
		threadsstr = strings.TrimSpace(threadsstr)
		threadsint, _ := strconv.Atoi(threadsstr)
		fmt.Print("TweetID: ")
		fmt.Scanln(&tweet_id)
		handlers.HandleRT(tokens, tweet_id, threadsint)
		goto restart
	case "6":
		var prompt string
		var countstr string
		var threadsstr string
		fmt.Print("Count: ")
		fmt.Scanln(&countstr)
		fmt.Print("Threads: ")
		fmt.Scanln(&threadsstr)
		fmt.Print("Prompt: ")
		fmt.Scanln(&prompt)

		countstr = strings.TrimSpace(countstr) // Trim leading and trailing whitespaces, including newline
		threadsstr = strings.TrimSpace(threadsstr)
		countint, err := strconv.Atoi(countstr)
		threadsint, _ := strconv.Atoi(threadsstr)
		if err != nil {
			fmt.Println("Error converting count to integer:", err)
			fmt.Println("Please enter a valid integer for count.")
			return
		}

		handlers.AiHandle(prompt, countint, threadsint)
		goto restart
	default:
		fmt.Println("Invalid module.")
		goto restart
	}
}
