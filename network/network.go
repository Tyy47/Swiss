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

func Connection() {
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		utils.Crash(err)
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		utils.Crash(err)
	}

	defer conn.Close()
	fmt.Println(status)
}

func GetPortStatus() {
	port, err := strconv.Atoi(port)
	if err != nil || port < 1 || port > 65535 {
		utils.Error("Port " + strconv.Itoa(port) + " is CLOSED on " + endpoint + ".")
		utils.Reason("Port exceeds or is under port range 0-65535")
		return
	}

	address := endpoint + ":" + strconv.Itoa(port)
	timeout := 3 * time.Second

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		utils.Error("Port " + strconv.Itoa(port) + " is CLOSED on " + endpoint + ".")
		utils.Reason(err.Error())
		return
	}

	conn.Close()

	utils.Success("Port " + strconv.Itoa(port) + " is OPEN on " + endpoint + ".")
}

// Takes an endpoint as a string and prints the IPv4 and v6 address of the domain.
func GetAddresses() {
	conn, err := net.LookupIP(endpoint)
	networkCrashError(err, "IP addresses")

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
