/*
Copyright AppsCode Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/avct/uasurfer"
	"github.com/go-macaron/auth"
	"github.com/go-macaron/binding"
	"github.com/go-macaron/cache"
	"github.com/gocarina/gocsv"
	csvtypes "gomodules.xyz/encoding/csv/types"
	freshsalesclient "gomodules.xyz/freshsales-client-go"
	gdrive "gomodules.xyz/gdrive-utils"
	listmonkclient "gomodules.xyz/listmonk-client-go"
	"gomodules.xyz/sets"
	"gopkg.in/macaron.v1"
)

type SpeakerInfo struct {
	Name     string `json:"name"`
	JobTitle string `json:"job_title"`
	Picture  string `json:"picture"`
}

type WebinarSchedule struct {
	Title          string         `json:"title" csv:"Title" form:"title"`
	Schedules      csvtypes.Dates `json:"schedules" csv:"Schedules" form:"schedules"`
	Summary        string         `json:"summary" csv:"Summary" form:"summary"`
	Speaker        string         `json:"speaker" csv:"Speaker" form:"speaker"`
	SpeakerTitle   string         `json:"speaker_title" csv:"Speaker Title" form:"speaker_title"`
	SpeakerBio     string         `json:"speaker_bio" csv:"Speaker Bio" form:"speaker_bio"`
	SpeakerPicture string         `json:"speaker_picture" csv:"Speaker Picture" form:"speaker_picture"`
	Speakers       []SpeakerInfo  `json:"speakers" csv:"-" form:"-"`
	YoutubeLink    string         `json:"-" csv:"Youtube Link" form:"-"`
	YoutubeVideoID string         `json:"youtube_video_id" csv:"-" form:"-"`
}

type WebinarMeetingID struct {
	GoogleCalendarEventID string `json:"google_calendar_event_id" csv:"Google Calendar Event ID"`
	ZoomMeetingID         int    `json:"zoom_meeting_id" csv:"Zoom Meeting ID"`
	ZoomMeetingPassword   string `json:"zoom_meeting_password" csv:"Zoom Meeting Password"`
}

type WebinarInfo struct {
	WebinarSchedule
	WebinarMeetingID
}

type WebinarRegistrationForm struct {
	Schedule string `json:"schedule" csv:"-" form:"schedule"`

	FirstName string `json:"first_name" csv:"First Name" form:"first_name"`
	LastName  string `json:"last_name" csv:"Last Name" form:"last_name"`
	Phone     string `json:"phone" csv:"Phone" form:"phone"`
	JobTitle  string `json:"job_title" csv:"Job Title" form:"job_title"`
	Company   string `json:"company" csv:"Company" form:"company"`
	WorkEmail string `json:"work_email" csv:"Work Email" form:"work_email"`

	ClusterProvider csvtypes.StringSlice `json:"cluster_provider,omitempty" csv:"Cluster Provider" form:"cluster_provider"`
	ExperienceLevel string               `json:"experience_level,omitempty" csv:"Experience Level" form:"experience_level"`
	MarketingReach  string               `json:"marketing_reach,omitempty" csv:"Marketing Reach" form:"marketing_reach"`
}

type WebinarRegistrationEmail struct {
	WorkEmail string `json:"work_email" csv:"Work Email" form:"work_email"`
}

func (s *Server) RegisterWebinarAPI(m *macaron.Macaron) {
	m.Get("/_/webinars", func(ctx *macaron.Context, c cache.Cache, log *log.Logger) {
		key := ctx.Req.URL.String()
		out := c.Get(key)
		if out == nil {
			schedule, err := s.NextWebinarSchedule()
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			out = schedule
			_ = c.Put(key, out, 600) // cache for 10 mins
		} else {
			log.Println(key, "found")
		}
		ctx.JSON(http.StatusOK, out)
	})

	m.Get("/_/upcoming_webinars", func(ctx *macaron.Context, c cache.Cache, log *log.Logger) {
		key := ctx.Req.URL.String()
		out := c.Get(key)
		if out == nil {
			schedule, err := s.UpcomingWebinarSchedules()
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			out = schedule
			_ = c.Put(key, out, 600) // cache for 10 mins
		} else {
			log.Println(key, "found")
		}
		ctx.JSON(http.StatusOK, out)
	})

	m.Get("/_/past_webinars", func(ctx *macaron.Context, c cache.Cache, log *log.Logger) {
		key := ctx.Req.URL.String()
		out := c.Get(key)
		if out == nil {
			schedule, err := s.PastWebinarSchedules()
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			out = schedule
			_ = c.Put(key, out, 600) // cache for 10 mins
		} else {
			log.Println(key, "found")
		}
		ctx.JSON(http.StatusOK, out)
	})

	m.Post("/_/webinars/register", binding.Bind(WebinarRegistrationForm{}), func(ctx *macaron.Context, form WebinarRegistrationForm, log *log.Logger) {
		err := s.RegisterForWebinar(ctx, form, log)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.WriteHeader(http.StatusOK)
		// ctx.JSON(http.StatusOK, form)
		// ctx.Redirect("https://appscode.com", http.StatusSeeOther)
	})

	m.Get("/_/webinars/:date/emails", auth.Basic(os.Getenv("APPSCODE_PRICING_USERNAME"), os.Getenv("APPSCODE_PRICING_PASSWORD")), func(ctx *macaron.Context) {
		date := ctx.Params("date")
		attendees, err := s.ListWebinarAttendees(date)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, attendees)
	})
}

func (s *Server) ListWebinarAttendees(date string) ([]string, error) {
	reader, err := gdrive.NewColumnReader(s.srvSheets, WebinarSpreadsheetId, date, "Work Email")
	if err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	attendees := []*WebinarRegistrationEmail{}
	if err := gocsv.UnmarshalCSV(reader, &attendees); err != nil { // Load attendees from file
		return nil, err
	}

	result := sets.NewString()
	for _, v := range attendees {
		result.Insert(v.WorkEmail)
	}
	return result.List(), nil
}

func (s *Server) NextWebinarSchedule() (*WebinarSchedule, error) {
	now := time.Now()

	reader, err := gdrive.NewRowReader(s.srvSheets, WebinarSpreadsheetId, WebinarScheduleSheet, &gdrive.Predicate{
		Header: "Schedules",
		By: func(column []interface{}) (int, error) {
			for i, v := range column {
				schedules := csvtypes.Dates{}
				err := schedules.UnmarshalCSV(v.(string))
				if err != nil {
					return -1, err
				}

				// 3/11/2021 3:00:00
				for _, t := range schedules {
					if t.After(now) {
						return i, nil
					}
				}
			}
			return -1, io.EOF
		},
	})
	if err == io.EOF {
		return &WebinarSchedule{}, nil
	} else if err != nil {
		return nil, err
	}

	schedules := []*WebinarSchedule{}
	if err := gocsv.UnmarshalCSV(reader, &schedules); err != nil { // Load clients from file
		return nil, err
	}
	for i := 0; i < len(schedules); {
		sch := schedules[i]

		var dates []time.Time
		for _, t := range sch.Schedules {
			if t.After(now) {
				dates = append(dates, t)
			}
		}
		if len(dates) == 0 {
			// pass webinar
			schedules = append(schedules[:i], schedules[i+1:]...)
			continue
		}
		sch.Schedules = []time.Time{dates[0]}
		if videoID, err := YoutubeVideoID(sch.YoutubeLink); err != nil {
			return nil, err
		} else {
			sch.YoutubeVideoID = videoID
		}
		FixSpeakers(sch)
		i++
	}
	sort.Slice(schedules, func(i, j int) bool {
		return schedules[i].Schedules[0].Before(schedules[j].Schedules[0])
	})

	if len(schedules) > 0 {
		return schedules[0], nil
	}
	return &WebinarSchedule{}, nil
}

func (s *Server) UpcomingWebinarSchedules() ([]*WebinarSchedule, error) {
	now := time.Now()

	reader, err := gdrive.NewRowReader(s.srvSheets, WebinarSpreadsheetId, WebinarScheduleSheet, &gdrive.Predicate{
		Header: "Schedules",
		By: func(column []interface{}) (int, error) {
			pos := math.MaxInt
			for i, v := range column {
				schedules := csvtypes.Dates{}
				err := schedules.UnmarshalCSV(v.(string))
				if err != nil {
					return -1, err
				}

				// 3/11/2021 3:00:00
				for _, t := range schedules {
					if t.After(now) && i < pos {
						pos = i
					}
				}
			}
			if pos == math.MaxInt {
				return -1, io.EOF
			}
			return pos, nil
		},
	})
	if err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	schedules := []*WebinarSchedule{}
	if err := gocsv.UnmarshalCSV(reader, &schedules); err != nil { // Load clients from file
		return nil, err
	}

	for i := 0; i < len(schedules); {
		sch := schedules[i]

		var dates []time.Time
		for _, t := range sch.Schedules {
			if t.After(now) {
				dates = append(dates, t)
			}
		}
		if len(dates) == 0 {
			// pass webinar
			schedules = append(schedules[:i], schedules[i+1:]...)
			continue
		}
		sch.Schedules = []time.Time{dates[0]}
		if videoID, err := YoutubeVideoID(sch.YoutubeLink); err != nil {
			return nil, err
		} else {
			sch.YoutubeVideoID = videoID
		}
		FixSpeakers(sch)
		i++
	}
	sort.Slice(schedules, func(i, j int) bool {
		return schedules[i].Schedules[0].Before(schedules[j].Schedules[0])
	})
	return schedules, nil
}

func (s *Server) PastWebinarSchedules() ([]*WebinarSchedule, error) {
	now := time.Now()

	reader, err := gdrive.NewReader(s.srvSheets, WebinarSpreadsheetId, WebinarScheduleSheet, 1)
	if err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	schedules := []*WebinarSchedule{}
	if err := gocsv.UnmarshalCSV(reader, &schedules); err != nil { // Load clients from file
		return nil, err
	}

	for i := 0; i < len(schedules); {
		sch := schedules[i]

		var dates []time.Time
		for _, t := range sch.Schedules {
			if t.Before(now) {
				dates = append(dates, t)
			}
		}
		if len(dates) == 0 {
			// pass webinar
			schedules = append(schedules[:i], schedules[i+1:]...)
			continue
		}
		sch.Schedules = []time.Time{dates[0]}
		if videoID, err := YoutubeVideoID(sch.YoutubeLink); err != nil {
			return nil, err
		} else {
			sch.YoutubeVideoID = videoID
		}
		FixSpeakers(sch)
		i++
	}
	sort.Slice(schedules, func(i, j int) bool {
		return schedules[i].Schedules[0].Before(schedules[j].Schedules[0])
	})
	return schedules, nil
}

/*
https://www.youtube.com/embed/_rVS3oe1usA
https://www.youtube.com/watch?v=_rVS3oe1usA
https://youtu.be/tjHvdkwSDF4
*/
func YoutubeVideoID(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	if id := u.Query().Get("v"); id != "" {
		return id, nil
	}
	return path.Base(u.Path), nil
}

func FixSpeakers(sch *WebinarSchedule) []SpeakerInfo {
	spliter := func(r rune) bool {
		return r == ',' || r == ';'
	}
	names := strings.FieldsFunc(sch.Speaker, spliter)
	titles := strings.FieldsFunc(sch.SpeakerTitle, spliter)
	pics := strings.FieldsFunc(sch.SpeakerPicture, spliter)

	n := min(len(names), min(len(titles), len(pics)))
	out := make([]SpeakerInfo, n)
	for i := 0; i < n; i++ {
		out[i] = SpeakerInfo{
			Name:     names[i],
			JobTitle: titles[i],
			Picture:  pics[i],
		}
	}

	// FIX input, if more than one speaker, set old fields for the first speaker
	sch.Speakers = out
	if n > 1 {
		sch.Speaker = out[0].Name
		sch.SpeakerTitle = out[0].JobTitle
		sch.SpeakerPicture = out[0].Picture
	}

	return out
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (s *Server) RegisterForWebinar(ctx *macaron.Context, form WebinarRegistrationForm, log *log.Logger) error {
	schedule, err := time.Parse(time.RFC3339, form.Schedule)
	if err != nil {
		return err
	}
	sheetName := schedule.Format("2006-1-2")
	clients := []*WebinarRegistrationForm{
		&form,
	}
	writer := gdrive.NewWriter(s.srvSheets, WebinarSpreadsheetId, sheetName)
	err = gocsv.MarshalCSV(clients, writer)
	if err != nil {
		return err
	}

	// create zoom, google calendar event if not exists,
	// add attendant if google calendar meeting exists

	// These api calls take too long to front proxies like Cloudflare to think server is unresponsive.
	// So, we return as soon as attendee name is recorded in a the Google spreadsheet.

	// nolint:errcheck
	go func() {
		err := func() error {
			yw, mw, dw := schedule.Date()

			reader, err := gdrive.NewRowReader(s.srvSheets, WebinarSpreadsheetId, "Schedule", &gdrive.Predicate{
				Header: "Schedules",
				By: func(values []interface{}) (int, error) {
					for i, v := range values {
						schedules := csvtypes.Dates{}
						err := schedules.UnmarshalCSV(v.(string))
						if err != nil {
							return -1, err
						}
						for _, t2 := range schedules {
							y2, m2, d2 := t2.Date()

							if yw == y2 && mw == m2 && dw == d2 {
								return i, nil
							}
						}
					}
					return -1, io.EOF
				},
			})
			if err != nil {
				return err
			}

			meetings := []*WebinarInfo{}
			if err := gocsv.UnmarshalCSV(reader, &meetings); err != nil { // Load clients from file
				return err
			}

			var result *WebinarInfo
			if len(meetings) > 0 {
				result = meetings[0]
			}
			if result == nil {
				return fmt.Errorf("can't find webinar schedule")
			}

			{
				// record in listmonk
				ml, err := s.listmonk.CreateListIfMissing(listmonkclient.MailingListRequest{
					Name:  fmt.Sprintf("webinar-%s", sheetName),
					Type:  listmonkclient.ListTypePrivate,
					Optin: listmonkclient.OptinModeSingle,
					Tags: []string{
						"webinar",
					},
				})
				if err != nil {
					return err
				}
				err = s.listmonk.SubscribeToList(listmonkclient.SubscribeRequest{
					Email:        form.WorkEmail,
					Name:         form.FirstName + " " + form.LastName,
					MailingLists: []string{ml.UUID},
				})
				if err != nil {
					return err
				}
			}

			{
				// record in CRM
				ua := uasurfer.Parse(ctx.Req.UserAgent())
				location := GeoLocation{
					IP: GetIP(ctx.Req.Request),
				}
				DecorateGeoData(s.geodb, &location)

				_ = s.noteEventWebinarRegistration(form, EventWebinarRegistration{
					BaseNoteDescription: freshsalesclient.BaseNoteDescription{
						Event: "webinar_registration",
						Client: freshsalesclient.ClientInfo{
							OS:     ua.OS.Name.StringTrimPrefix(),
							Device: ua.DeviceType.StringTrimPrefix(),
							Location: freshsalesclient.GeoLocation{
								IP:          location.IP,
								Timezone:    location.Timezone,
								City:        location.City,
								Country:     location.Country,
								Coordinates: location.Coordinates,
							},
						},
					},
					Webinar: WebinarRecord{
						Title:           result.Title,
						Schedule:        csvtypes.Timestamp{Time: schedule},
						Speaker:         result.Speaker,
						ClusterProvider: form.ClusterProvider,
						ExperienceLevel: form.ExperienceLevel,
						MarketingReach:  form.MarketingReach,
					},
				})
			}

			if strings.TrimSpace(result.GoogleCalendarEventID) != "" {
				wats, err := gdrive.NewColumnReader(s.srvSheets, WebinarSpreadsheetId, sheetName, "Work Email")
				if err != nil {
					return err
				}
				atts := []*WebinarRegistrationEmail{}
				if err := gocsv.UnmarshalCSV(wats, &atts); err != nil { // Load clients from file
					return err
				}

				emails := make([]string, len(atts))
				for i, a := range atts {
					emails[i] = a.WorkEmail
				}
				return AddEventAttendants(s.srvCalendar, WebinarCalendarId, result.GoogleCalendarEventID, emails)
			}

			ww := gdrive.NewRowWriter(s.srvSheets, WebinarSpreadsheetId, "Schedule", &gdrive.Predicate{
				Header: "Schedules",
				By: func(values []interface{}) (int, error) {
					for i, v := range values {
						schedules := csvtypes.Dates{}
						err := schedules.UnmarshalCSV(v.(string))
						if err != nil {
							return -1, err
						}
						for _, t2 := range schedules {
							y2, m2, d2 := t2.Date()

							if yw == y2 && mw == m2 && dw == d2 {
								return i, nil
							}
						}
					}
					return -1, io.EOF
				},
			})

			meetinginfo, err := CreateZoomMeeting(s.srvCalendar, s.zc, WebinarCalendarId, s.zoomAccountEmail, &result.WebinarSchedule, schedule, 60*time.Minute, []string{
				form.WorkEmail,
			})
			if err != nil {
				return err
			}

			meetings2 := []*WebinarMeetingID{
				meetinginfo,
			}
			return gocsv.MarshalCSV(meetings2, ww)
		}()
		if err != nil {
			log.Printf("failed to register for request: %+v, reason: %v", form, err)
		}
	}()

	return nil
}
