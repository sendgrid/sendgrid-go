package taskrouter

import (
	"fmt"
	"strings"
)

const (
	Version          = "v1"
	RouterBaseUrl    = "https://taskrouter.twilio.com"
	WebsocketBaseUrl = "https://event-bridge.twilio.com/v1/wschannels"
	Get              = "GET"
	Post             = "POST"
)

func Workspaces() string {
	return fmt.Sprint(strings.Join([]string{RouterBaseUrl, Version, "Workspaces"}, "/"))
}

func Workspace(workspaceSid string) string {
	return fmt.Sprint(strings.Join([]string{RouterBaseUrl, Version, "Workspaces", workspaceSid}, "/"))
}

func AllWorkspaces() string {
	return fmt.Sprint(strings.Join([]string{RouterBaseUrl, Version, "Workspaces", "**"}, "/"))
}

func Tasks(workspaceSid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "Tasks"}, "/"))
}

func Task(workspaceSid string, tasksSid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "Tasks", tasksSid}, "/"))
}

func AllTasks(workspaceSid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "Tasks", "**"}, "/"))
}

func TaskQueues(workspaceSid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "TaskQueues"}, "/"))
}

func TaskQueue(workspaceSid string, taskQueueSid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "TaskQueues", taskQueueSid}, "/"))
}

func AllTaskQueues(workspaceSid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "TaskQueues", "**"}, "/"))
}

func Activities(workspaceSid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "Activities"}, "/"))
}

func Activity(workspaceSid string, activitySid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "Activities", activitySid}, "/"))
}

func AllActivities(workspaceSid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "Activities", "**"}, "/"))
}

func Workers(workspaceSid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "Workers"}, "/"))
}

func Worker(workspaceSid string, workerSid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "Workers", workerSid}, "/"))
}

func AllWorkers(workspaceSid string) string {
	return fmt.Sprint(strings.Join([]string{Workspace(workspaceSid), "Workers", "**"}, "/"))
}

func Reservations(workspaceSid string, workerSid string) string {
	return fmt.Sprint(strings.Join([]string{Worker(workspaceSid, workerSid), "Reservations"}, "/"))
}

func Reservation(workspaceSid string, workerSid string, reservationSid string) string {
	return fmt.Sprint(strings.Join([]string{Worker(workspaceSid, workerSid), "Reservations", reservationSid}, "/"))
}

func AllReservations(workspaceSid string, workerSid string) string {
	return fmt.Sprint(strings.Join([]string{Worker(workspaceSid, workerSid), "Reservations", "**"}, "/"))
}

func WebSocketPolicies(accountSid string, channelSid string) []Policy {
	url := fmt.Sprint(strings.Join([]string{WebsocketBaseUrl, accountSid, channelSid}, "/"))
	get := GeneratePolicy(url, Get, true, nil, nil)
	post := GeneratePolicy(url, Post, true, nil, nil)
	return []Policy{get, post}
}

func WorkerPolicies(workspaceSid string, workerSid string) []Policy {
	activities := GeneratePolicy(Activities(workspaceSid), Get, true, nil, nil)
	tasks := GeneratePolicy(AllTasks(workspaceSid), Get, true, nil, nil)
	reservations := GeneratePolicy(AllReservations(workspaceSid, workerSid), Get, true, nil, nil)
	fetch := GeneratePolicy(Worker(workspaceSid, workerSid), Get, true, nil, nil)
	return []Policy{activities, tasks, reservations, fetch}
}
