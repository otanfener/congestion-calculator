package service

import (
	"context"
	"github.com/otanfener/congestion-controller/pkg/domain"
	"github.com/otanfener/congestion-controller/pkg/models"
	"github.com/otanfener/congestion-controller/repos"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	repo repos.Repo
}

func New(repo repos.Repo) *Service {
	srv := &Service{
		repo: repo,
	}
	return srv
}

func (srv *Service) CalculateTax(ctx context.Context, dates []models.CivilTime, city string, vehicle string) (models.Tax, error) {

	if dates == nil {
		return models.Tax{
			TotalFee: 0,
			History:  nil,
		}, nil
	}
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j].Time)
	})
	chargeHistoryPerDate := make(map[string]int)
	c, err := srv.repo.GetCity(ctx, city)
	if err != nil {
		return models.Tax{}, domain.ErrNotFound
	}

	if isTollFreeVehicle(vehicle, c.ExemptVehicles) {
		return models.Tax{TotalFee: 0, History: nil}, nil
	}
	var filteredDates []models.CivilTime

	for _, d := range dates {
		if isTollFreeDate(d.Time, c.Rules) {
			continue
		}
		filteredDates = append(filteredDates, d)
	}
	chargesPerDay := getTollFeeBySingleCharge(filteredDates, c.Rules, c.Tariff)

	total := calculateTotalTaxBySingleChargeRule(chargeHistoryPerDate, c.Rules, chargesPerDay)
	return models.Tax{
		TotalFee: total,
		History:  chargeHistoryPerDate,
	}, nil
}
func calculateTotalTaxBySingleChargeRule(chargeHistoryPerDay map[string]int, rules models.Rules, chargesPerDay map[string][]int) int {
	totalFee := 0
	for date, charges := range chargesPerDay {
		totalChargePerDay := 0
		for _, charge := range charges {
			totalChargePerDay += charge
		}
		if totalChargePerDay >= rules.MaxChargePerDay {
			totalChargePerDay = rules.MaxChargePerDay
		}
		chargeHistoryPerDay[date] = totalChargePerDay
		totalFee += totalChargePerDay
	}
	return totalFee
}
func getTollFeeBySingleCharge(dates []models.CivilTime, rules models.Rules, tariffs []models.Tariff) map[string][]int {
	visited := make([]time.Time, 0)
	result := make(map[string][]int)

	for start := 0; start < len(dates); start++ {
		if isVisited(visited, dates[start].Time) {
			continue
		}
		charge := getTollFeeByTariffAndDate(dates[start], tariffs)
		for end := start + 1; end < len(dates); end++ {
			duration := dates[end].Sub(dates[start].Time)
			durationInMin := int(duration.Minutes())
			if durationInMin <= rules.SingleChargeInterval {
				temp := getTollFeeByTariffAndDate(dates[end], tariffs)
				visited = append(visited, dates[end].Time)
				if temp >= charge {
					charge = temp
				}
			} else {
				break
			}
		}
		constructChargesByDate(dates, result, start, charge)
	}
	return result
}
func constructChargesByDate(dates []models.CivilTime, result map[string][]int, index int, charge int) {
	date := dates[index].Format("2006-01-02")
	if _, ok := result[date]; !ok {
		result[date] = make([]int, len(dates))
	}
	result[date][index] = charge
}
func isVisited(times []time.Time, t time.Time) bool {
	for _, tt := range times {
		if tt.Equal(t) {
			return true
		}
	}
	return false
}
func getTollFeeByTariffAndDate(date models.CivilTime, tariffs []models.Tariff) int {
	totalFee := 0
	if len(tariffs) == 0 {
		return totalFee
	}

	for _, tariff := range tariffs {
		hourFrom, _ := strconv.Atoi(strings.Split(tariff.From, ":")[0])
		minuteFrom, _ := strconv.Atoi(strings.Split(tariff.From, ":")[1])
		hourTo, _ := strconv.Atoi(strings.Split(tariff.To, ":")[0])
		minuteTo, _ := strconv.Atoi(strings.Split(tariff.To, ":")[1])

		fromTime := time.Date(date.Year(), date.Month(), date.Day(), hourFrom, minuteFrom, 0, date.Nanosecond()-1, date.Location())
		toTime := time.Date(date.Year(), date.Month(), date.Day(), hourTo, minuteTo, 59, date.Nanosecond(), date.Location())

		if isInTimeRange(date.Time, fromTime, toTime) {
			totalFee += tariff.Fee
			return totalFee
		}
	}
	return totalFee
}
func isInTimeRange(date time.Time, start time.Time, end time.Time) bool {
	return date.After(start) && date.Before(end)
}

func isTollFreeVehicle(vehicle string, exemptVehicles []string) bool {
	if len(exemptVehicles) == 0 {
		return false
	}
	for _, e := range exemptVehicles {
		if strings.EqualFold(vehicle, e) {
			return true
		}
	}
	return false
}

func isTollFreeDate(date time.Time, rules models.Rules) bool {

	for _, h := range rules.OfficialHolidays {
		d, _ := time.Parse("2006-01-02", h)
		dateWithTimeZone := time.Date(d.Year(), d.Month(), d.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
		if dateWithTimeZone.Equal(date) || dateWithTimeZone.AddDate(0, 0, -1).Equal(date) {
			return true
		}
	}
	for _, m := range rules.ChargeFreeMonths {
		if strings.EqualFold(date.Month().String(), strings.ToLower(m)) {
			return true
		}
	}
	for _, w := range rules.WorkingDays {
		if strings.EqualFold(date.Weekday().String(), strings.ToLower(w)) {
			return false
		}
	}

	return true
}
