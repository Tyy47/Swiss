package network

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

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

func Connection(endpoint string) {
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	fmt.Println(status)
}

// Takes an endpoint as a string and prints the IPv4 and v6 address of the domain.
func GetAddresses(endpoint string, writeToFile bool) {
	conn, err := net.LookupIP(endpoint)
	if err != nil {
		utils.Error("Unable to connect: Unknown host.")
	}

	if writeToFile {
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
	} else {
		for rows := range conn {
			fmt.Println(conn[rows])
		}
	}
}

func GetNameServer(endpoint string, writeToFile bool) {
	conn, err := net.LookupNS(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	if writeToFile {
		write := bufio.NewWriter(outputFile.file)
		if _, err := write.WriteString("\nName Servers: \n"); err != nil {
			utils.CrashCheck(err)
		}
		for rows := range conn {
			val := conn[rows].Host + "\n"

			_, err := write.WriteString(val)
			if err != nil {
				log.Fatal(err)
			}
		}

		write.Flush()
	} else {
		utils.Success("\nNameservers for " + endpoint + ".")
		for rows := range conn {
			val := conn[rows].Host + "\n"
			fmt.Println(val)
		}
	}
}

func GetCNameRecords(endpoint string, writeToFile bool) {
	conn, err := net.LookupCNAME(endpoint)
	utils.CrashCheck(err)

	if writeToFile {
		write := bufio.NewWriter(outputFile.file)
		if _, err := write.WriteString("\nCName Records: \n"); err != nil {
			utils.CrashCheck(err)
		}

		_, err := write.WriteString(conn)
		if err != nil {
			utils.CrashCheck(err)
		}

		write.Flush()
	} else {
		utils.Success("CName Records for " + endpoint + ".")
		fmt.Println(conn)
	}
}

func GetTXTRecords(endpoint string, writeToFile bool) {
	conn, err := net.LookupTXT(endpoint)
	utils.CrashCheck(err)

	if writeToFile {
		write := bufio.NewWriter(outputFile.file)
		if _, err := write.WriteString("\n\nTXT Records: \n"); err != nil {
			utils.CrashCheck(err)
		}

		for rows := range conn {
			_, err := write.WriteString(conn[rows] + "\n")
			utils.CrashCheck(err)
		}

		write.Flush()
	} else {
		utils.Success("TXT Records for " + endpoint + ".")
		for rows := range conn {
			fmt.Println(conn[rows])
		}
	}
}

func GetMXRecords(endpoint string, writeToFile bool) {
	conn, err := net.LookupMX(endpoint)
	utils.CrashCheck(err)

	if writeToFile {
		write := bufio.NewWriter(outputFile.file)
		if _, err := write.WriteString("\nMX Records: \n"); err != nil {
			utils.CrashCheck(err)
		}

		for rows := range conn {
			_, err := write.WriteString("Host: " + conn[rows].Host + "\n")
			utils.CrashCheck(err)
		}

		write.Flush()
	} else {
		utils.Success("MX Records for " + endpoint + ".")

		for rows := range conn {
			fmt.Println("Host: " + conn[rows].Host + "\n")
		}
	}
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
func GatherData(endpoint string) {
	initFileCreation()
	GetAddresses(endpoint, true)
	GetNameServer(endpoint, true)
	GetCNameRecords(endpoint, true)
	GetTXTRecords(endpoint, true)
	GetMXRecords(endpoint, true)
	utils.Success("Written data to " + outputFile.name + " for endpoint " + endpoint + ".")
}
