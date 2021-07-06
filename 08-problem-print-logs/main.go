package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Log entry
type log struct {
	method    string
	status    string
	endpoint  string
	hit_count int32
}

// Map of log entries
var logs map[string]*log

// Slice of sorted (by hits count) logs map keys in descending order
var sortedLogKeys []string

// List of active users
var activeUsers map[int]bool

// Get log entries
// Split logs by new line character (\n), write
// individual log strings to output channel
// and return output channel
func getLogEntries(logs string) <-chan string {
	out := make(chan string)

	// Write individual log entries to output channel
	go func(out chan<- string) {
		defer close(out)
		logsSlice := strings.Split(logs, "\n") // Split logs by newline character (\n)

		for _, log := range logsSlice {
			out <- log // Write to output channel
		}
	}(out)

	return out
}

// Replace numeric userId in URI with a character
func replaceUserIdsWithChar(endpoint string, replaceWith string) string {
	// numRegex := regexp.MustCompile(`\d`)
	numRegex := regexp.MustCompile(`[0-9]+`)
	return numRegex.ReplaceAllString(endpoint, replaceWith)
}

// Build key from "method", "endpoint" and "status_code"
func getLogKey(logFields []string) string {
	endpoint := replaceUserIdsWithChar(logFields[2], "#")
	key := logFields[1] + endpoint + logFields[3]
	return strings.ToLower(key)
}

// Extract userId from endpoint URI
func getUserId(endpoint string) int {
	fields := strings.Split(strings.ToLower(endpoint), "/")

	if len(fields) > 2 && strings.Contains(fields[1], "user") {
		uid, err := strconv.Atoi(fields[2])
		if err != nil {
			fmt.Println(err)
			return 0
		}
		return uid
	}
	return 0
}

// Process independent log entries and build logs map
func processLogs(in <-chan string) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	logs = make(map[string]*log)
	activeUsers = make(map[int]bool)

	for logEntry := range in {
		logFields := strings.Split(logEntry, " ")
		if len(logFields) >= 4 {
			wg.Add(1)

			go func() {
				defer wg.Done()
				defer mu.Unlock()
				mu.Lock()

				userId := getUserId(logFields[2])
				if userId > 0 {
					activeUsers[userId] = true
				}

				key := getLogKey(logFields)
				if _, ok := logs[key]; ok {
					logs[key].hit_count++
					return
				} else {
					logs[key] = &log{
						method:    logFields[1],
						status:    logFields[3],
						endpoint:  replaceUserIdsWithChar(logFields[2], "#"),
						hit_count: 1,
					}
				}
			}()
		}
	}

	wg.Wait()
}

// Sort logs by hits count
func sortLogsByHits() {
	for key := range logs {
		sortedLogKeys = append(sortedLogKeys, key)
	}
	sort.Slice(sortedLogKeys, func(i, j int) bool {
		return logs[sortedLogKeys[i]].hit_count > logs[sortedLogKeys[j]].hit_count
	})
}

// Print output
func printOutput() {
	fmt.Printf("%-10s %-20s %-10s %-10s\n", "Method", "Endpoint", "Status", "Hits")
	fmt.Println("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
	// for _, log := range logs {
	// 	fmt.Printf("%-10s %-20s %-10s %-10d\n", log.method, log.endpoint, log.status, log.hit_count)
	// }
	for _, logKey := range sortedLogKeys {
		fmt.Printf("%-10s %-20s %-10s %-10d\n", logs[logKey].method, logs[logKey].endpoint, logs[logKey].status, logs[logKey].hit_count)
	}
	fmt.Println("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")

	actUsers := "Active Users: "
	for uid := range activeUsers {
		actUsers += fmt.Sprintf("%d, ", uid)
	}
	fmt.Println(actUsers[:len(actUsers)-2])
}

func main() {
	logs := "[01-02-2021:08:22:32] GET /users/12 200\n" +
		"[01-03-2021:08:22:32] GET /users/12 200\n" +
		"[01-03-2021:09:22:32] GET /users/13 200\n" +
		"[22-03-2021:01:22:32] POST /users/15 200\n" +
		"[22-03-2021:01:22:32] GET /users/55/friends 200\n" +
		"[22-03-2021:01:22:32] DELETE /users/64 200\n" +
		"[10-03-2021:01:22:32] GET /cart 200\n" +
		"-------SERVER DOWN-------\n" +
		"[02-03-2021:01:22:32] GET /cart 200\n" +
		"[11-03-2021:01:22:32] POST /users/16 200\n" +
		"[01-03-2021:01:22:32] POST /users 201\n" +
		"[02-03-2021:01:22:32] GET /orders 200\n" +
		"[20-03-2021:01:22:32] POST /logout 200\n"

	processLogs(getLogEntries(logs))
	sortLogsByHits()
	printOutput()
}

// Output:
/*
Method     Endpoint             Status     Hits
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
GET        /users/#             200        3
POST       /users/#             200        2
GET        /cart                200        2
GET        /orders              200        1
POST       /users               201        1
DELETE     /users/#             200        1
POST       /logout              200        1
GET        /users/#/friends     200        1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
Active Users: 64, 12, 15, 55, 13, 16
*/
