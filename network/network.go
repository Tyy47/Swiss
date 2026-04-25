package network

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"swiss/utils"
)

type FileReturn struct {
	name string
	file *os.File
	err  string
}

var outputFile = FileReturn{
	name: "swiss_net_output",
	err:  "File already exists",
}

var (
	endpoint = utils.CheckArguments(utils.Arguments, 3, 3)
	port     = utils.CheckArguments(utils.Arguments, 4, 4)
)

func networkCrashError(err error, data string) {
	if err != nil {
		utils.Error("Unable to gather " + data + " from " + endpoint + ".")
		os.Exit(1)
	}
}

// Connects to an IP Address or domain via TCP and returns an HTTP status code that states the result of the connection
func Connection() {
	// Connects to the domain or IP address
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		utils.Crash(err)
	}
	
	// Prints the status code neatly
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		utils.Crash(err)
	}
	
	// Closes the connection when everything is complete
	defer conn.Close()
	// Prints the result of the http status code.
	fmt.Println(status)
}

// Connects to an IP Address or Domain through a port, prints a statement regarding the status of the port.
func GetPortStatus() {
	// Grabs the port from the package arguments and checks if it is valid.
	port, err := strconv.Atoi(port)
	if err != nil || port < 1 || port > 65535 {
		utils.Error("Port " + strconv.Itoa(port) + " is CLOSED on " + endpoint + ".")
		utils.Reason("Port exceeds or is under port range 0-65535")
		return
	}
	
	// Concats the address into a full string with the port
	address := endpoint + ":" + strconv.Itoa(port)
	// Timeout configured for the port.
	timeout := 3 * time.Second
	
	// Connects to the address with the newly given port
	conn, err := net.DialTimeout("tcp", address, timeout)

	// Handles an error if the connection fails
	if err != nil {
		utils.Error("Port " + strconv.Itoa(port) + " is CLOSED on " + endpoint + ".")
		utils.Reason(err.Error())
		return
	}
	
	// Closes the connection when the connection attempt is completed
	defer conn.Close()

	// Success message if the port is open on the address.
	utils.Success("Port " + strconv.Itoa(port) + " is OPEN on " + endpoint + ".")
}

// Takes an endpoint as a string and prints the IPv4 and v6 address of the domain.
func GetAddresses() {
	// Connects to the domain via a given address
	conn, err := net.LookupIP(endpoint)
	// Crashes the program if there is a network connection related error.
	networkCrashError(err, "IP addresses") 
	
	// Prints all IP addresses that are tied to that address to the console.
	for rows := range conn {
		fmt.Println(conn[rows])
	}
}

func writeAddresses() {
	conn, err := net.LookupIP(endpoint)
	networkCrashError(err, "IP Addresses")
	write := bufio.NewWriter(outputFile.file)

	if _, err := write.WriteString("\nIP Addresses: \n"); err != nil {
		utils.CrashCheck(err)
	}

	for rows := range conn {
		val := conn[rows].String()

		if _, err := write.WriteString(val + "\n"); err != nil {
			utils.CrashCheck(err)
		}
	}

	write.Flush()
}

func GetNameServer() {
	conn, err := net.LookupNS(endpoint)
	networkCrashError(err, "Name servers")

	utils.Success("\nNameservers for " + endpoint + ".")
	for rows := range conn {
		val := conn[rows].Host + "\n"
		fmt.Println(val)
	}
}

func writeNameServer() {
	conn, err := net.LookupNS(endpoint)
	networkCrashError(err, "Name servers")

	write := bufio.NewWriter(outputFile.file)
	if _, err := write.WriteString("\nName Servers: \n"); err != nil {
		utils.CrashCheck(err)
	}
	for rows := range conn {
		val := conn[rows].Host + "\n"

		_, err := write.WriteString(val)
		if err != nil {
			utils.Crash(err)
		}
	}

	write.Flush()
}

func GetCNameRecords() {
	conn, err := net.LookupCNAME(endpoint)
	networkCrashError(err, "CNAME records")

	utils.Success("CNAME Records for " + endpoint + ".")
	fmt.Println(conn)
}

func writeCNameRecords() {
	conn, err := net.LookupCNAME(endpoint)
	networkCrashError(err, "CNAME records")

	write := bufio.NewWriter(outputFile.file)
	if _, err := write.WriteString("\nCNAME Records: \n"); err != nil {
		utils.CrashCheck(err)
	}

	_, errr := write.WriteString(conn)
	if errr != nil {
		utils.CrashCheck(err)
	}

	write.Flush()
}

func GetTXTRecords() {
	conn, err := net.LookupTXT(endpoint)
	networkCrashError(err, "TXT records")

	utils.Success("TXT Records for " + endpoint + ".")
	for rows := range conn {
		fmt.Println(conn[rows])
	}
}

func writeTXTRecords() {
	conn, err := net.LookupTXT(endpoint)
	networkCrashError(err, "TXT records")

	write := bufio.NewWriter(outputFile.file)
	if _, err := write.WriteString("\n\nTXT Records: \n"); err != nil {
		utils.CrashCheck(err)
	}

	for rows := range conn {
		_, err := write.WriteString(conn[rows] + "\n")
		utils.CrashCheck(err)
	}

	write.Flush()
}

func GetMXRecords() {
	conn, err := net.LookupMX(endpoint)
	networkCrashError(err, "MX records")

	utils.Success("MX Records for " + endpoint + ".")

	for rows := range conn {
		fmt.Println("Host: " + conn[rows].Host + "\n")
	}
}

func writeMXRecords() {
	conn, err := net.LookupMX(endpoint)
	networkCrashError(err, "MX records")

	write := bufio.NewWriter(outputFile.file)
	if _, err := write.WriteString("\nMX Records: \n"); err != nil {
		utils.CrashCheck(err)
	}

	for rows := range conn {
		_, err := write.WriteString("Host: " + conn[rows].Host + "\n")
		utils.CrashCheck(err)
	}

	write.Flush()
}

// Creates the output file for net
func initFileCreation() {
	if !utils.CheckFileExists(outputFile.name) {
		utils.MakeFile(outputFile.name, true)
		file, err := os.Create(outputFile.name)
		utils.CrashCheck(err)
		outputFile.file = file
	} else {
		// removes the old data
		os.Remove(outputFile.name)
		utils.MakeFile(outputFile.name, true)
		file, err := os.Create(outputFile.name)
		utils.CrashCheck(err)
		outputFile.file = file
	}
}

// Culminates all of the data in the above functions, then outputs all the data into a file for viewing.
func GatherData() {
	initFileCreation()
	writeAddresses()
	writeNameServer()
	writeCNameRecords()
	writeTXTRecords()
	writeMXRecords()
	utils.Success("Written data to " + outputFile.name + " for endpoint " + endpoint + ".")
}
