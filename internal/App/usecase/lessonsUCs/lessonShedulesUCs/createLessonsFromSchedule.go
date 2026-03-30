package lessonshedulesucs

import (
	"context"
	"time"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	usecaseutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/usecase_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type CreateLessonsFromShceduleUC struct {
	repoLessons   entitiesrepos.LessonsRepo
	repoSchedules entitiesrepos.ScheduleRepo
}

type ICreateLessonsFromShceduleUC interface {
	Execute(ctx context.Context) error
}

func NewCreateLessonsFromShceduleUC(
	repo entitiesrepos.LessonsRepo,
	repoSchedules entitiesrepos.ScheduleRepo) ICreateLessonsFromShceduleUC {
	return &CreateLessonsFromShceduleUC{repo, repoSchedules}
}

func (uc *CreateLessonsFromShceduleUC) Execute(ctx context.Context) error {
	schedules := []entities.Schedule{}
	if err := uc.repoSchedules.GetAll(ctx, &schedules); err != nil {
		return err
	}

	today := time.Now()
	limit := today.AddDate(0, 1, 0) // + 1 mounth

	lessonsToInsert := make([]entities.Lesson, 0, len(schedules)*5+10)

	for i := 0; i < len(schedules); i++ {
		schedule := schedules[i]

		start := usecaseutils.MaxTime(schedule.Active_from, today)
		end := usecaseutils.MinTime(schedule.Active_to, limit)

		if schedule.Active_to.Before(today) {
			continue
		} else if start.After(end) {
			continue
		}

		dates := usecaseutils.GetDatesByWeekday(
			start,
			end,
			time.Weekday(schedule.Weekday),
		)

		for d := 0; d < len(dates); d++ {
			hour, min, sec := 0, 0, 0
			if schedule.Start_time != "" {
				t, err := time.Parse("15:04:05", schedule.Start_time)
				if err == nil {
					hour, min, sec = t.Hour(), t.Minute(), t.Second()
				}
			}
			lessonsToInsert = append(lessonsToInsert, entities.Lesson{
				Group_id:     schedule.Group_id,
				Direction_id: schedule.Direction_id,
				Teacher_id:   schedule.Teacher_id,
				Lesson_date: time.Date(
					dates[d].Year(),
					dates[d].Month(),
					dates[d].Day(),
					hour,
					min,
					sec,
					0,
					dates[d].Location(),
				),
				Duration_minutes: schedule.Duration_minutes,
				Is_canceled:      false,
				School_id:        schedule.School_id,
			})
		}
	}

	if len(lessonsToInsert) == 0 {
		return nil
	}

	return uc.repoLessons.CreateMany(ctx, &lessonsToInsert)
}
