package activities

import (
	"maps"
	"slices"
)

type Activity struct {
	Distance  int
	EventTime int64
	TimeStamp int64
}

type Solution struct {
	users map[string]map[string]Activity
	// TODO: Add a field for storing scheduled events
}

func NewSolution() *Solution {
	return &Solution{
		users: make(map[string]map[string]Activity),
		// TODO: Initialize the field for scheduled events
	}
}

func (s *Solution) AddActivity(userId, activityType string, distance int) bool {
	if _, exists := s.users[userId]; !exists {
		s.users[userId] = make(map[string]Activity)
	}

	activity := s.users[userId][activityType]
	if distance < 0 || (activity.Distance > 0 && activity.EventTime == 0) {
		return false
	}

	// TODO: Integrate scheduled events distance here
	activity = Activity{
		Distance:  distance + activity.Distance,
		EventTime: 0,
		TimeStamp: 0,
	}

	s.users[userId][activityType] = activity
	return true
}

func (s *Solution) UpdateActivity(userId, activityType string, distance int) bool {
	if userActivities, exists := s.users[userId]; exists {
		if _, exists := userActivities[activityType]; exists {

			activity := Activity{
				Distance:  s.users[userId][activityType].Distance,
				EventTime: s.users[userId][activityType].EventTime,
				TimeStamp: s.users[userId][activityType].TimeStamp,
			}
			activity.Distance += distance

			s.users[userId][activityType] = activity
			return true
		}
	}
	return false
}

func (s *Solution) GetActivity(userId, activityType string) *int {
	if userActivities, exists := s.users[userId]; exists {
		if activity, exists := userActivities[activityType]; exists {
			return &activity.Distance
		}
	}
	return nil
}

func (s *Solution) ActivitySummary(userId string) map[string]int {
	userActivities := s.users[userId]

	summary := make(map[string]int)
	for activityType, activity := range userActivities {
		summary[activityType] = activity.Distance
	}

	return summary
}

// TODO: Implement the ScheduleEvent function
func (s *Solution) ScheduleEvent(timestamp int64, userId, activityType string, distance int, eventTime int64) bool {
	if timestamp < 0 || eventTime < 0 || eventTime < timestamp {
		return false
	}

	if _, exists := s.users[userId]; !exists {
		s.users[userId] = make(map[string]Activity)
	}

	activity := Activity{
		Distance:  distance,
		EventTime: eventTime,
		TimeStamp: timestamp,
	}

	s.users[userId][activityType] = activity
	return true
}

// TODO: Implement the GetAgenda function
func (s *Solution) GetAgenda(userId string, fromTime, toTime int64) []map[string]interface{} {
	if _, exists := s.users[userId]; !exists {
		return make([]map[string]interface{}, 0)
	}

	agenda := make([]map[string]interface{}, 0)
	userActivities := s.users[userId]
	for _, key := range slices.Sorted(maps.Keys(userActivities)) {
		activity := userActivities[key]
		if activity.EventTime == 0 {
			continue
		}

		if activity.EventTime >= fromTime && activity.EventTime <= toTime {
			agenda = append(agenda, map[string]interface{}{
				"activityType": key,
				"distance":     activity.Distance,
				"eventTime":    activity.EventTime,
			})
		}
	}

	return agenda
}
